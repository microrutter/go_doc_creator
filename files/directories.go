package files

import (
	"io/fs"
	"log"
	"path/filepath"

	"github.com/microrutter/go_doc_creator/utils"
)

func NewDirectories() *Directories {
	return &Directories{}
}

func NewDirectory() *Directory {
	return &Directory{}
}

func NewFile() *File {
	return &File{}
}

func (nd *Directories) WalkAllFilesInDir(dir string, log *log.Logger, confpath string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, e error) error {
		utils.Check(e, log)
		if info.Mode().IsRegular() {
			log.Printf("Found new file at path %s", path)
			doc := NewDocument()
			doc.ReadFile(log, path, confpath)
			dirName := filepath.Base(filepath.Dir(path))
			dd := NewDirectory()
			dd.Name = dirName
			nf := NewFile()
			nf.Name = info.Name()
			nf.Doc = *doc
			dd.Files = append(dd.Files, *nf)
			nd.ListDirect = append(nd.ListDirect, *dd)
		}
		return nil
	})
}
