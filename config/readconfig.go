package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"

	"github.com/microrutter/go_doc_creator/utils"
)

func Config() *Yaml {
	return &Yaml{}
}

func (c *Yaml) GetConf(log *log.Logger, filepath string) {
	log.Print("Getting Configuration")
	yamlFile, err := ioutil.ReadFile(filepath)
	utils.Check(err, log)
	err = yaml.Unmarshal(yamlFile, c)
	utils.Check(err, log)
}
