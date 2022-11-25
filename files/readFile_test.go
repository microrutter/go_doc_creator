package files

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/microrutter/go_doc_creator/config"
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
		text        string
		findText    config.Comment
		newLine     bool
		wantText    string
		wantNewLine bool
	}{
		{"//This is a comment", config.Comment{Start: "//"}, false, "This is a comment", false},
		{"// This is a comment", config.Comment{Start: "//"}, false, " This is a comment", false},
		{"#This is a comment", config.Comment{Start: "#"}, false, "This is a comment", false},
		{"https://api.pagerduty.com/incidents", config.Comment{Start: "//"}, false, "", false},
		{"//This is a comment", config.Comment{Start: "//", Finish: "//"}, false, "This is a comment", true},
		{"// This is a comment", config.Comment{Start: "//", Finish: "//"}, false, " This is a comment", true},
		{"#This is a comment", config.Comment{Start: "#", Finish: "//"}, false, "This is a comment", true},
		{"https://api.pagerduty.com/incidents", config.Comment{Start: "//", Finish: "//"}, false, "", false},
		{"This is a comment//", config.Comment{Start: "//", Finish: "//"}, true, "This is a comment", false},
		{" This is a comment //", config.Comment{Start: "//", Finish: "//"}, true, " This is a comment ", false},
		{"This is a comment#", config.Comment{Start: "#", Finish: "#"}, true, "This is a comment", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Test Name: %s", tt.text)
		t.Run(testname, func(t *testing.T) {
			ans, newl := comments(tt.text, tt.findText, tt.newLine)
			if ans != tt.wantText {
				t.Errorf("got %s, want %s", ans, tt.wantText)
			}

			if newl != tt.wantNewLine {
				t.Errorf("Expected result for newline not correct for test %s got %t", tt.text, newl)
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

func TestReadFileMultiple(t *testing.T) {
	secondCommentsTest := []string{"constants", " The following is going to Test", " this this and this good luck", " Set Constants", "More Constants for the before each"}
	secondSubtitleComments := []string{" We are going top do this", " Then we are going to do this", "  The we might do this"}
	secondSubTitle := []string{"First test woohoo", "Second test woohoo", "Third test woohoo", "Fourth test woohoo", "Fith test woohoo", "Sixth test woohoo"}
	r := NewDocument()
	r.ReadFile(logger, "../test_files/test_cypress_multiple_describes.cy.ts", "../test_files/test_conf.yaml")

	if len(r.Test) != 2 {
		t.Errorf("Expecting a length of 2 got %d", len(r.Test))
	}

	secondMainTitle := r.GetMainTitle(logger)
	secondMainComments := r.GetMainComments(logger)
	secondSubTitles := r.GetSubTitle(logger)

	if strings.Trim(secondMainTitle, " ") != "This is an amazing description tahnkl you" {
		t.Errorf("Second title was not correct looking for `This is an amazing description tahnkl you` got %s", strings.Trim(secondMainTitle, " "))
	}

	if reflect.DeepEqual(secondMainComments, secondCommentsTest) {
		t.Errorf("Second comments Array not equal")
	}

	if len(secondSubTitles) != 6 {
		t.Errorf("Not enough Subtitles for second test found")
	}

	for i, s := range secondSubTitles {
		if reflect.DeepEqual(s.Comment, secondSubtitleComments) {
			t.Errorf("secondSubTitle: %s does not have matching comments", s.Title)
		}

		if strings.Trim(s.Title, " ") != secondSubTitle[i] {
			t.Errorf("secondSubTitle was not correct looking for `%s` got %s", secondSubTitle[i], strings.Trim(s.Title, " "))
		}
	}

	mainCommentsTest := []string{"constants (Main)", " The following is going to Test (Main)", " this this and this good luck (Main)", " Set Constants (Main)", "More Constants for the before each (Main)"}
	subtitleComments := []string{" We are going top do this (Main)", " Then we are going to do this (Main)", "  The we might do this (Main)"}
	subTitle := []string{"First test woohoo (Main)", "Second test woohoo (Main)", "Third test woohoo (Main)", "Fourth test woohoo (Main)", "Fith test woohoo (Main)", "Sixth test woohoo (Main)"}

	mainTitle := r.Test[0].Title.Title
	mainComments := r.Test[0].Title.Comment
	subTitles := r.Test[0].SubTitle

	if strings.Trim(mainTitle, " ") != "This is an amazing description thank you I am also the main test" {
		t.Errorf("Main title was not correct looking for `This is an amazing description thank you I am also the main test` got %s", strings.Trim(mainTitle, " "))
	}

	if reflect.DeepEqual(mainComments, mainCommentsTest) {
		t.Errorf("Comments Array not equal")
	}

	if len(subTitles) != 6 {
		t.Errorf("Not enough Subtitles found expected 6 got %d", len(subTitles))
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
