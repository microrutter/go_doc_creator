package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/microrutter/go_doc_creator/notion"
)

func main() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
	)

	notion.CreateNotionDocuments(logger, "conf.yaml", "/media/wayne/FreeAgent GoFlex Drive/Central Data Store/plandek - wip/nextgen/e2e/cypress/e2e")

	fmt.Print(&buf)
}
