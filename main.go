package main

import (
	"bytes"
	"flag"
	"log"

	"github.com/microrutter/go_doc_creator/notion"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "logger: ", log.Lshortfile)
)

func main() {
	configfile := flag.String("configfile", "conf.yaml", "The file path to the conf file")
	filepath := flag.String("topdir", "/path/to/top/level/file", "The file path to where your cypress tests are held")
	flag.Parse()
	notion.CreateNotionDocuments(logger, *configfile, *filepath)

	// fmt.Print(&buf)
}
