package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/microrutter/go_doc_creator/files"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	d := files.NewDocument()

	d.ReadFile(*logger, "/media/wayne/FreeAgent GoFlex Drive/Central Data Store/plandek - wip/nextgen/e2e/cypress/e2e/data-integrity-int/pagerduty_elastic.cy.ts")

	logger.Printf("Main Doc Title: %s", d.GetMainTitle(*logger))

	fmt.Print(&buf)
}