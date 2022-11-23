package config

type Configuration struct {
	MainTitle  string `yaml:"maintitle"`
	TitleSplit Splits `yaml:"split"`
	SubTitle   string `yaml:"subtitle"`
	Comment    string `yaml:"comment"`
	Output     Output `yaml:"output"`
}

type Yaml struct {
	Conf Configuration `yaml:"conf"`
}

type Splits struct {
	Start  string `yaml:"start"`
	Finish string `yaml:"finish"`
}

type Output struct {
	Type         string `yaml:"type"`
	Secret       string `yaml:"secret"`
	StartingPage string `yaml:"startingpage"`
}
