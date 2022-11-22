package config

type Configuration struct {
	MainTitle  string `yaml:"maintitle"`
	TitleSplit Splits `yaml:"split"`
	SubTitle   string `yaml:"subtitle"`
	Comment    string `yaml:"comment"`
}

type Yaml struct {
	Conf Configuration `yaml:"conf"`
}

type Splits struct {
	Start  string `yaml:"start"`
	Finish string `yaml:"finish"`
}
