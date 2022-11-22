package files

type MainTitle struct {
	Title   string
	Comment []string
}

type SubTitles struct {
	Title   string
	Comment []string
}

type Document struct {
	Title    MainTitle
	SubTitle []SubTitles
}
