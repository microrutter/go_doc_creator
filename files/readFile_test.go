package files

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func TestTitle(t *testing.T) {
	var tests = []struct {
		text        string
		newLine     bool
		findText    string
		start       string
		stop        string
		wantText    string
		wantNewLine bool
	}{
		{"describe(\"This is a test\", function ())", false, "describe", "\"", "\"", "This is a test", false},
		{"describe(", false, "describe", "\"", "\"", "", true},
		{"\"This is the start", true, "describe", "\"", "\"", "This is the start", true},
		{"This is the finish\", function()", true, "describe", "\"", "\"", "This is the finish", false},
		{"it(\"This is a test\", function ())", false, "it(", "\"", "\"", "This is a test", false},
		{"it(\"This is a test looking for describe (\", function ())", false, "describe(", "\"", "\"", "", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test Name: %s", tt.text)
		t.Run(testname, func(t *testing.T) {
			ans, newl := title(tt.text, tt.newLine, tt.findText, tt.start, tt.stop)
			if ans != tt.wantText {
				t.Errorf("got %s, want %s", ans, tt.wantText)
			}

			if newl != tt.wantNewLine {
				t.Errorf("Wanted: %s, Got: %s", strconv.FormatBool(tt.wantNewLine), strconv.FormatBool(newl))
			}
		})
	}

}

func TestComments(t *testing.T) {
	var tests = []struct {
		text     string
		findText string
		wantText string
	}{
		{"//This is a comment", "//", "This is a comment"},
		{"// This is a comment", "//", " This is a comment"},
		{"#This is a comment", "#", "This is a comment"},
		{"https://api.pagerduty.com/incidents", "//", ""},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test Name: %s", tt.text)
		t.Run(testname, func(t *testing.T) {
			ans := comments(tt.text, tt.findText)
			if ans != tt.wantText {
				t.Errorf("got %s, want %s", ans, tt.wantText)
			}
		})
	}

}

func TestReadFile(t *testing.T) {
	mainCommentsTest := []string{"constants", " The following is going to Test", " this this and this good luck", " Set Constants", "More Constants for the before each"}
	subtitleComments := []string{" We are going top do this", " Then we are going to do this", "  The we might do this"}
	subTitle := []string{"First test woohoo", "Second test woohoo", "Third test woohoo", "Fourth test woohoo", "Fith test woohoo", "Sixth test woohoo"}
	r := NewDocument()
	r.ReadFile(logger, "../test_files/test_cypress_files.cy.ts", "../test_files/test_conf.yaml")

	fmt.Print(r.Title.Title)

	mainTitle := r.GetMainTitle(logger)
	mainComments := r.GetMainComments(logger)
	subTitles := r.GetSubTitle(logger)

	if strings.Trim(mainTitle, " ") != "This is an amazing description tahnkl you" {
		t.Errorf("Main title was not correct looking for `This is an amazing description tahnkl you` got %s", strings.Trim(mainTitle, " "))
	}

	if reflect.DeepEqual(mainComments, mainCommentsTest) {
		t.Errorf("Comments Array not equal")
	}

	if len(subTitles) != 6 {
		t.Errorf("Not enough Subtitles found")
	}

	for i, s := range subTitles {
		if reflect.DeepEqual(s.Comment, subtitleComments) {
			t.Errorf("SubTitle: %s does not have matching comments", s.Title)
		}

		if strings.Trim(s.Title, " ") != subTitle[i] {
			t.Errorf("SubTitle was not correct looking for `%s` got %s", subTitle[i], strings.Trim(s.Title, " "))
		}
	}
}
