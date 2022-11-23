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
	y.GetConf(logger, "../test_files/test_conf.yaml")

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
	if c.Output.Type != "test" {
		t.Errorf("Was looking for test but got %s", c.Output.Type)
	}
	if c.Output.Secret != "iamsecret" {
		t.Errorf("Was looking for iamsecret but got %s", c.Output.Secret)
	}
	if c.Output.StartingPage != "starthere" {
		t.Errorf("Was looking for starthere but got %s", c.Output.StartingPage)
	}
	if c.Output.Image != "iamaimageurl" {
		t.Errorf("was looking for iamaimageurl but got %s", c.Output.Image)
	}
}
