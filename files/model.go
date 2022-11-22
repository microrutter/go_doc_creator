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

type File struct {
	Name string
	Doc  Document
}

type Directory struct {
	Name  string
	Files []File
}

type Directories struct {
	ListDirect []Directory
}
