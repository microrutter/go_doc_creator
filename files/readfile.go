package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func NewDocument() *Document {
	return &Document{}
}

func newSubTitle() *SubTitles {
	return &SubTitles{}
}

func check(e error, log log.Logger) {
	if e != nil {
		log.Print(e)
	}
}

func (sub *SubTitles) setSubTitle(log log.Logger, title string) {
	log.Printf("Setting subtitle title to: %s", title)
	sub.Title = sub.Title + title + " "
}

func (sub *SubTitles) setSubTitleComments(log log.Logger, comment string) {
	log.Printf("Setting subtitle comment to: %s", comment)
	sub.Title = sub.Comment + comment + " "
}

func (newDoc *Document) GetSubTitle(log log.Logger) []SubTitles {
	log.Printf("Getting the list of subtitles")
	return newDoc.SubTitle
}

func (newDoc *Document) setMainTitle(log log.Logger, title string) {
	log.Printf("Setting Main Title To: %s", title)
	newDoc.Title.Title = newDoc.Title.Title + title + " "
}

func (newDoc *Document) GetMainTitle(log log.Logger) string {
	log.Printf("Getting Main Title")
	return newDoc.Title.Title
}

func (newDoc *Document) setMainComments(log log.Logger, comment string) {
	log.Printf("Setting Main Comment To: %s", comment)
	newDoc.Title.Title = newDoc.Title.Comment + comment + " "
}

func (newDoc *Document) GetMainComments(log log.Logger) string {
	log.Printf("Getting Main Comments")
	return newDoc.Title.Comment
}

func (newDoc *Document) addSubTitle(log log.Logger, subTitle SubTitles) {
	log.Printf("Adding a subtitle to main Document")
	newDoc.SubTitle = append(newDoc.SubTitle, subTitle)
}

func (newDoc *Document) GetLastSubTitle() *SubTitles {
	return &newDoc.SubTitle[len(newDoc.SubTitle)-1]
}

func (newDoc *Document) ReadFile(log log.Logger, filepath string) {

	log.Printf("Starting to read file at %s", filepath)

	f, err := os.Open(filepath)

	check(err, log)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	newLine := false

	for scanner.Scan() {
		log.Println("Checking Main Title")
		t, nextLine := title(scanner.Text(), newLine, "describe")
		newLine = nextLine
		newDoc.setMainTitle(log, t)
		if !newLine && len(newDoc.GetMainTitle(log)) > 0 && len(newDoc.GetSubTitle(log)) <= 0 {
			comment := comments(scanner.Text(), "//")
			newDoc.setMainComments(log, comment)
		}
		st, nextLine := title(scanner.Text(), newLine, "it(")
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
			comment := comments(scanner.Text(), "//")
			subTitle.setSubTitleComments(log, comment)
		}
	}
}

func title(text string, nextLine bool, findText string) (string, bool) {
	if strings.Contains(text, findText) || nextLine {
		var newS = ""
		s := strings.Index(text, "\"")
		if s == -1 || (strings.Count(text, "\"") == 1 && s != 0) {
			if !nextLine {
				return "", true
			}
			newS = text
		} else {
			newS = text[s+1:]
		}
		e := strings.Index(newS, "\"")
		if e == -1 {
			return newS, true
		}
		result := newS[:e]
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
