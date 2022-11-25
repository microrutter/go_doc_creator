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

func newTest() *Test {
	return &Test{}
}

func (sub *SubTitles) setSubTitle(log *log.Logger, title string) {
	log.Printf("Setting subtitle title to: %s", title)
	builder := strings.Builder{}
	builder.WriteString(sub.Title)
	builder.WriteString(title)
	builder.WriteString(" ")
	sub.Title = sub.Title + title + " "
}

func (sub *SubTitles) setSubTitleComments(log *log.Logger, comment string) {
	log.Printf("Setting subtitle comment to: %s", comment)
	sub.Comment = append(sub.Comment, comment)
}

func (newDoc *Document) GetSubTitle(log *log.Logger) []SubTitles {
	log.Printf("Getting the list of subtitles")
	return newDoc.GetMainTest(log).SubTitle
}

func (t *Test) setMainTitle(log *log.Logger, title string) {
	log.Printf("Setting Main Title To: %s", title)
	builder := strings.Builder{}
	builder.WriteString(t.Title.Title)
	builder.WriteString(title)
	builder.WriteString(" ")
	t.Title.Title = builder.String()
}

func (newDoc *Document) GetMainTitle(log *log.Logger) string {
	log.Printf("Getting Main Title")
	return newDoc.GetMainTest(log).Title.Title
}

func (newDoc *Document) GetMainTest(log *log.Logger) *Test {
	log.Printf("Getting Main Test")
	if len(newDoc.Test) > 0 {
		return &newDoc.Test[len(newDoc.Test)-1]
	}
	return &newDoc.Test[0]
}

func (newDoc *Document) setMainComments(log *log.Logger, comment string) {
	log.Printf("Setting Main Comment To: %s", comment)
	newDoc.GetMainTest(log).Title.Comment = append(newDoc.GetMainTest(log).Title.Comment, comment)
}

func (newDoc *Document) setNewTest(log *log.Logger, t Test) {
	log.Println("Adding new test to Document")
	newDoc.Test = append(newDoc.Test, t)
}

func (newDoc *Document) GetMainComments(log *log.Logger) []string {
	log.Printf("Getting Main Comments")
	return newDoc.GetMainTest(log).Title.Comment
}

func (newDoc *Document) addSubTitle(log *log.Logger, subTitle SubTitles) {
	log.Printf("Adding a subtitle to main Document")
	newDoc.GetMainTest(log).SubTitle = append(newDoc.GetMainTest(log).SubTitle, subTitle)
}

func (newDoc *Document) GetLastSubTitle(log *log.Logger) *SubTitles {
	return &newDoc.GetMainTest(log).SubTitle[len(newDoc.GetMainTest(log).SubTitle)-1]
}

func (newDoc *Document) ReadFile(log *log.Logger, filepath string, conffile string) error {

	yaml := config.Config()

	yaml.GetConf(log, conffile)

	conf := yaml.Conf

	log.Printf("Starting to read file at %s", filepath)

	f, err := os.Open(filepath)

	utils.Check(err, log)

	scanner := bufio.NewScanner(f)

	newLine := false

	commentCont := false

	for scanner.Scan() {
		log.Println("Checking Main Title")
		text := scanner.Text()
		t, nextLine := title(text, newLine, conf.MainTitle, conf.TitleSplit.Start, conf.TitleSplit.Finish)
		if len(strings.TrimSpace(t)) > 0 && !newLine {
			tt := newTest()
			tt.setMainTitle(log, t)
			newDoc.setNewTest(log, *tt)
		} else if len(strings.TrimSpace(t)) > 0 && newLine {
			newDoc.GetMainTest(log).setMainTitle(log, t)
		}
		newLine = nextLine
		if !newLine && len(newDoc.Test) > 0 && len(newDoc.GetSubTitle(log)) <= 0 {
			comment, cont := comments(text, conf.Comment, commentCont)
			newDoc.setMainComments(log, comment)
			commentCont = cont
		}
		st, nextLine := title(text, newLine, conf.SubTitle, conf.TitleSplit.Start, conf.TitleSplit.Finish)
		if len(st) > 0 && !nextLine {
			subTitle := newSubTitle()
			subTitle.setSubTitle(log, st)
			newDoc.addSubTitle(log, *subTitle)
		}

		if len(st) > 0 && nextLine {
			subTitle := newDoc.GetLastSubTitle(log)
			subTitle.setSubTitle(log, st)
			newDoc.addSubTitle(log, *subTitle)
		}

		if len(newDoc.Test) > 0 && len(newDoc.GetSubTitle(log)) > 0 {
			subTitle := newDoc.GetLastSubTitle(log)
			comment, cont := comments(text, conf.Comment, commentCont)
			subTitle.setSubTitleComments(log, comment)
			commentCont = cont
		}

	}

	return f.Close()
}

func title(text string, nextLine bool, findText string, start string, stop string) (string, bool) {
	if strings.Contains(strings.TrimSpace(text), findText) && strings.Index(strings.TrimSpace(text), findText) < 4 || nextLine {
		return getRequiredText(text, start, stop, nextLine)
	}
	return "", false
}

func comments(text string, findText config.Comment, nextLine bool) (string, bool) {
	compareText := strings.TrimSpace(text)
	if strings.Contains(compareText, findText.Start) && strings.Index(compareText, findText.Start) < 4 || nextLine {
		if len(findText.Finish) == 0 {
			return strings.TrimPrefix(compareText, findText.Start), false
		} else {
			return getRequiredText(text, findText.Start, findText.Finish, nextLine)
		}
	}
	return "", false
}

func getRequiredText(text string, start string, stop string, nextLine bool) (string, bool) {
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
