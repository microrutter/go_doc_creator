package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func check(e error, log log.Logger) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFile(log log.Logger, filepath string) Document {

	var newDoc Document

	log.Printf("Starting to read file at %s", filepath)

	f, err := os.Open(filepath)

	check(err, log)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	newLine := false

	skipTitle := false

	for scanner.Scan() {
		log.Println("Checking Main Title")
		t, nextLine, skip := title(scanner.Text(), newLine, skipTitle)
		if (skip && !skipTitle) || nextLine {
			skipTitle = skip
			newLine = nextLine
			newDoc.Title.Title = newDoc.Title.Title + t
			log.Printf("Title: %s", newDoc.Title.Title)
		}
	}
	return newDoc
}

func title(text string, nextLine bool, skip bool) (string, bool, bool) {
	if (strings.Contains(text, "describe") || nextLine) && !skip {
		var newS = ""
		s := strings.Index(text, "\"")
		if s == -1 || (strings.Count(text, "\"") == 1 && s != 0) {
			if !nextLine {
				return "", true, false
			}
			newS = text
		} else {
			newS = text[s+1:]
		}
		e := strings.Index(newS, "\"")
		if e == -1 {
			return newS, true, false
		}
		result := newS[:e]
		return result, false, true
	}
	return "", false, true
}
