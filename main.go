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

	d := files.NewDirectories()

	d.WalkAllFilesInDir("/media/wayne/FreeAgent GoFlex Drive/Central Data Store/plandek - wip/nextgen/e2e/cypress/e2e", *logger, "conf.yaml")

	fmt.Print(&buf)
}
