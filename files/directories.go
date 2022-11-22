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

func (nd *Directories) WalkAllFilesInDir(dir string, log log.Logger, confpath string) error {
	return filepath.Walk(dir, func(path string, info fs.FileInfo, e error) error {
		utils.Check(e, &log)
		if info.Mode().IsRegular() {
			log.Printf("Found new file at path %s", path)
			nf := NewFile()
			nf.Name = info.Name()
			log.Println(nf.Name)
			dirName := filepath.Base(filepath.Dir(path))
			dd := checkIfDirectoryExists(nd, dirName, log)
			doc := NewDocument()
			doc.ReadFile(&log, path, confpath)
			nf.Doc = *doc
			dd.Files = append(dd.Files, *nf)
			nd.ListDirect = append(nd.ListDirect, *dd)
		}

		return nil
	})
}

func checkIfDirectoryExists(d *Directories, dn string, log log.Logger) *Directory {
	for _, dir := range d.ListDirect {
		if dir.Name == dn {
			log.Printf("found %s", dn)
			return &dir
		}
	}
	log.Printf("Not found %s creating new directory", dn)
	nd := NewDirectory()
	nd.Name = dn
	return nd
}
