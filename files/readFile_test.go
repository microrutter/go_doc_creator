package files

import (
	"bytes"
	"log"
	"testing"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func TestReadFile(t *testing.T) {
	text, newLine, skip := title("describe(\"This is a test\", function ())", false, false)

	if newLine {
		t.Error("NewLine came back as true")
	}

	if !skip {
		t.Error("Skip came back as False")
	}

	if text != "This is a test" {
		t.Errorf("Text cam back as: %s", text)
	}

	nlText, nlnewLine, nlskip := title("describe(", false, false)

	if !nlnewLine {
		t.Error("NewLine came back as false")
	}

	if nlskip {
		t.Error("Skip came back as True")
	}

	if nlText != "" {
		t.Errorf("Text cam back as: %s", nlText)
	}

	ptext, pnewLine, pskip := title("\"This is the start", true, false)

	if !pnewLine {
		t.Error("NewLine came back as false")
	}

	if pskip {
		t.Error("Skip came back as True")
	}

	if ptext != "This is the start" {
		t.Errorf("Text came back as: %s", ptext)
	}

	lasttext, lastnewLine, lastskip := title("This is the finish\", function()", true, false)

	if lastnewLine {
		t.Error("Last NewLine came back as true")
	}

	if !lastskip {
		t.Error("Last Skip came back as false")
	}

	if lasttext != "This is the finish" {
		t.Errorf("Last Text came back as: %s", lasttext)
	}
}
