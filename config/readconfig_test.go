package config

import (
	"bytes"
	"log"
	"testing"
)

var (
	buf    bytes.Buffer
	logger = log.New(&buf, "test_logger: ", log.Lshortfile)
)

func TestGetConfig(t *testing.T) {
	y := Config()
	y.GetConf(*logger, "../test_files/test_conf.yaml")

	c := y.Conf

	if c.MainTitle != "describe" {
		t.Errorf("Was looking for describe but got %s", c.MainTitle)
	}
	if c.TitleSplit.Start != "\"" {
		t.Errorf("Was looking for \" but got %s", c.TitleSplit.Start)
	}
	if c.TitleSplit.Finish != "\"" {
		t.Errorf("Was looking for \" but got %s", c.TitleSplit.Finish)
	}
	if c.SubTitle != "it(" {
		t.Errorf("Was looking for it( but got %s", c.SubTitle)
	}
	if c.Comment != "//" {
		t.Errorf("Was looking for // but got %s", c.Comment)
	}
}
