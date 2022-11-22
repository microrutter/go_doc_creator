package files

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/microrutter/go_doc_creator/config"
	"github.com/microrutter/go_doc_creator/utils"
)

func NewDocument() *Document {
	return &Document{}
}

func newSubTitle() *SubTitles {
	return &SubTitles{}
}

func (sub *SubTitles) setSubTitle(log *log.Logger, title string) {
	log.Printf("Setting subtitle title to: %s", title)
	sub.Title = sub.Title + title + " "
}

func (sub *SubTitles) setSubTitleComments(log *log.Logger, comment string) {
	log.Printf("Setting subtitle comment to: %s", comment)
	sub.Title = sub.Comment + comment + " "
}

func (newDoc *Document) GetSubTitle(log *log.Logger) []SubTitles {
	log.Printf("Getting the list of subtitles")
	return newDoc.SubTitle
}

func (newDoc *Document) setMainTitle(log *log.Logger, title string) {
	log.Printf("Setting Main Title To: %s", title)
	newDoc.Title.Title = newDoc.Title.Title + title + " "
}

func (newDoc *Document) GetMainTitle(log *log.Logger) string {
	log.Printf("Getting Main Title")
	return newDoc.Title.Title
}

func (newDoc *Document) setMainComments(log *log.Logger, comment string) {
	log.Printf("Setting Main Comment To: %s", comment)
	newDoc.Title.Title = newDoc.Title.Comment + comment + " "
}

func (newDoc *Document) GetMainComments(log *log.Logger) string {
	log.Printf("Getting Main Comments")
	return newDoc.Title.Comment
}

func (newDoc *Document) addSubTitle(log *log.Logger, subTitle SubTitles) {
	log.Printf("Adding a subtitle to main Document")
	newDoc.SubTitle = append(newDoc.SubTitle, subTitle)
}

func (newDoc *Document) GetLastSubTitle() *SubTitles {
	return &newDoc.SubTitle[len(newDoc.SubTitle)-1]
}

func (newDoc *Document) ReadFile(log *log.Logger, filepath string, conffile string) error {

	yaml := config.Config()

	yaml.GetConf(*log, conffile)

	conf := yaml.Conf

	log.Printf("Starting to read file at %s", filepath)

	f, err := os.Open(filepath)

	utils.Check(err, log)

	scanner := bufio.NewScanner(f)

	newLine := false

	for scanner.Scan() {
		log.Println("Checking Main Title")
		text := scanner.Text()
		t, nextLine := title(text, newLine, conf.MainTitle, conf.TitleSplit.Start, conf.TitleSplit.Finish)
		newLine = nextLine
		newDoc.setMainTitle(log, t)
		if !newLine && len(newDoc.GetMainTitle(log)) > 0 && len(newDoc.GetSubTitle(log)) <= 0 {
			comment := comments(text, conf.Comment)
			newDoc.setMainComments(log, comment)
		}
		st, nextLine := title(text, newLine, conf.SubTitle, conf.TitleSplit.Start, conf.TitleSplit.Finish)
		if len(st) > 0 && !nextLine {
			subTitle := newSubTitle()
			subTitle.setSubTitle(log, st)
			newDoc.addSubTitle(log, *subTitle)
		}

		if len(st) > 0 && nextLine {
			subTitle := newDoc.GetLastSubTitle()
			subTitle.setSubTitle(log, st)
			newDoc.addSubTitle(log, *subTitle)
		}

		if len(newDoc.GetSubTitle(log)) > 0 {
			subTitle := newDoc.GetLastSubTitle()
			comment := comments(text, conf.Comment)
			subTitle.setSubTitleComments(log, comment)
		}
	}

	return f.Close()
}

func title(text string, nextLine bool, findText string, start string, stop string) (string, bool) {
	if strings.Contains(text, findText) || nextLine {
		var newS = ""
		s := strings.Index(text, start)
		if s == -1 || (strings.Count(text, stop) == 1 && s != 0) {
			if !nextLine {
				return "", true
			}
			newS = text
		} else {
			_, newS, _ = strings.Cut(text, start)
		}
		e := strings.Index(newS, stop)
		if e == -1 {
			return newS, true
		}
		result, _, _ := strings.Cut(newS, stop)
		return result, false
	}
	return "", false
}

func comments(text string, findText string) string {
	compareText := strings.Trim(text, " ")
	if strings.Contains(compareText, findText) && strings.Index(compareText, findText) < 4 {
		return strings.TrimPrefix(compareText, findText)
	}
	return ""
}
