package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/microrutter/go_doc_creator/config"
	"github.com/microrutter/go_doc_creator/files"
	"github.com/microrutter/go_doc_creator/utils"
)

var (
	baseUrl       = "https://api.notion.com/v1/"
	notionVersion = "2022-06-28"
)

func NewResult() *Result {
	return &Result{}
}

func NewBody() *Body {
	return &Body{}
}

func NewBodyNoChildren() *BodyNoChildren {
	return &BodyNoChildren{}
}

func NewText() *Text {
	return &Text{}
}

func NewContent() *Content {
	return &Content{}
}

func NewRichText() *RichText {
	return &RichText{}
}

func NewHeadingOneObject() *HeadingOneObject {
	return &HeadingOneObject{}
}

func NewHeadingTwoObject() *HeadingTwoObject {
	return &HeadingTwoObject{}
}

func NewBulletedListObject() *BulletedListObject {
	return &BulletedListObject{}
}

func addHeader(req *http.Request, secret string) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", secret))
	req.Header.Set("Notion-version", notionVersion)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func deletePages(log *log.Logger, conffile string, conf *config.Configuration) {
	url := fmt.Sprintf("blocks/%s/children", conf.Output.StartingPage)
	client, reqH := createReqClient(log, "GET", url, conf.Output.Secret, nil)
	resp, err := client.Do(reqH)
	utils.Check(err, log)
	out := NewResult()
	err = json.NewDecoder(resp.Body).Decode(&out)
	utils.Check(err, log)
	resp.Body.Close()
	for _, o := range out.Object {
		if o.Type == "child_page" {
			deleteUrl := fmt.Sprintf("blocks/%s", o.Id)
			delClient, delResp := createReqClient(log, "DELETE", deleteUrl, conf.Output.Secret, nil)
			respDel, err := delClient.Do(delResp)
			utils.Check(err, log)
			log.Printf("Response code: %d from deletion of %s", resp.StatusCode, o.Id)
			respDel.Body.Close()
		}
	}
}

func CreateNotionDocuments(log *log.Logger, conffile string, filepath string) {
	yaml := config.Config()
	yaml.GetConf(log, conffile)
	conf := yaml.Conf
	deletePages(log, conffile, &conf)
	rf := files.NewDirectories()
	rf.WalkAllFilesInDir(filepath, log, conffile)
	for _, dir := range rf.ListDirect {
		url := fmt.Sprintf("blocks/%s/children", conf.Output.StartingPage)
		cli, reqH := createReqClient(log, "GET", url, conf.Output.Secret, nil)
		oResp, err := cli.Do(reqH)
		utils.Check(err, log)
		oldOut := NewResult()
		err = json.NewDecoder(oResp.Body).Decode(&oldOut)
		utils.Check(err, log)
		oResp.Body.Close()
		pageExists := true
		for _, exist := range oldOut.Object {
			if exist.Type == "child_page" && exist.ChildPage.Title == dir.Name {
				id := exist.Id
				addFilePages(log, id, dir, conf.Output.Secret, conf.Output.Image)
				pageExists = false
			}
		}
		if pageExists {
			bd := NewBodyNoChildren()
			bd.Parent.PageId = conf.Output.StartingPage
			bd.Cover.External.URL = conf.Output.Image
			tt := NewText()
			tt.Text.Cont = dir.Name
			bd.Properties.Text = append(bd.Properties.Text, *tt)
			sendBody, err := json.Marshal(bd)
			utils.Check(err, log)
			responseBody := bytes.NewBuffer(sendBody)
			client, req := createReqClient(log, "POST", "pages", conf.Output.Secret, responseBody)
			resp, err := client.Do(req)
			utils.Check(err, log)
			log.Printf("Response code: %d from creation of %s", resp.StatusCode, dir.Name)
			out := NewResult()
			err = json.NewDecoder(resp.Body).Decode(&out)
			utils.Check(err, log)
			resp.Body.Close()
			id := out.Id
			addFilePages(log, id, dir, conf.Output.Secret, conf.Output.Image)
		}

	}
}

func createReqClient(log *log.Logger, op string, url string, secret string, body io.Reader) (http.Client, *http.Request) {
	log.Printf("Creating request and client for url: %s with op: %s", url, op)
	client := http.Client{}
	req, err := http.NewRequest(op, fmt.Sprintf("%s%s", baseUrl, url), body)
	utils.Check(err, log)
	reqH := addHeader(req, secret)
	return client, reqH
}

func addFilePages(log *log.Logger, id string, dir files.Directory, secret string, imageurl string) {

	for _, file := range dir.Files {
		subPage := NewBody()
		subPage.Parent.PageId = id
		subPage.Cover.External.URL = imageurl
		ntt := NewText()
		ntt.Text.Cont = file.Name
		subPage.Properties.Text = append(subPage.Properties.Text, *ntt)
		ho := NewHeadingOneObject()
		ho.addHeadingOne(log, file.Doc.Title.Title)
		subPage.Children = append(subPage.Children, *ho)

		for _, s := range file.Doc.Title.Comment {
			if len(strings.TrimSpace(s)) > 0 {
				b := NewBulletedListObject()
				b.addComments(log, s)
				subPage.Children = append(subPage.Children, *b)
			}
		}

		for _, sp := range file.Doc.SubTitle {
			ht := NewHeadingTwoObject()
			ht.addHeadingTwo(log, sp.Title)
			subPage.Children = append(subPage.Children, *ht)
			for _, s := range sp.Comment {
				if len(strings.TrimSpace(s)) > 0 {
					sb := NewBulletedListObject()
					sb.addComments(log, s)
					subPage.Children = append(subPage.Children, *sb)
				}
			}
		}
		sendBody, err := json.Marshal(subPage)
		utils.Check(err, log)
		responseBody := bytes.NewBuffer(sendBody)
		client, req := createReqClient(log, "POST", "pages", secret, responseBody)
		resp, err := client.Do(req)
		utils.Check(err, log)
		log.Printf("Response code: %d from creation of %s", resp.StatusCode, file.Name)
		resp.Body.Close()
	}
}

func (b *BulletedListObject) addComments(log *log.Logger, comments string) {
	b.Object = "block"
	b.Type = "bulleted_list_item"
	rt := NewRichText()
	nntt := NewContent()
	nntt.Cont = strings.TrimSpace(comments)
	rt.Type = "text"
	rt.Text = *nntt
	b.BulletList.RichText = append(b.BulletList.RichText, *rt)
}

func (h *HeadingOneObject) addHeadingOne(log *log.Logger, title string) {
	h.Object = "block"
	h.Type = "heading_1"
	rt := NewRichText()
	nntt := NewContent()
	nntt.Cont = strings.TrimSpace(title)
	rt.Type = "text"
	rt.Text = *nntt
	h.HeadOne.RichText = append(h.HeadOne.RichText, *rt)
}

func (h *HeadingTwoObject) addHeadingTwo(log *log.Logger, title string) {
	h.Object = "block"
	h.Type = "heading_2"
	rt := NewRichText()
	nntt := NewContent()
	nntt.Cont = strings.TrimSpace(title)
	rt.Type = "text"
	rt.Text = *nntt
	h.HeadTwo.RichText = append(h.HeadTwo.RichText, *rt)
}
