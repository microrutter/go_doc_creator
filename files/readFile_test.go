package files

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
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
	r := NewDocument()
	r.ReadFile(logger, "../test_files/test_cypress_files.cy.ts", "../test_files/test_conf.yaml")

	fmt.Print(r.Title.Title)

	mainTitle := r.GetMainTitle(logger)

	if mainTitle != "This is an amazing description tahnkl you" {
		t.Errorf("Main title was not correct looking for `This is an amazing description tahnkl you` got %s", mainTitle)
	}
	fmt.Print(&buf)
}
