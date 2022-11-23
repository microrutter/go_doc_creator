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
			// dd, appendDoc := checkIfDirectoryExists(*nd, dirName, log)
			nf, appendFile := checkIfFileExists(dd.Files, info.Name(), log)
			nf.Doc = *doc
			if appendFile {
				dd.Files = append(dd.Files, *nf)
			}
			nd.ListDirect = append(nd.ListDirect, *dd)
		}

		return nil
	})
}

func checkIfDirectoryExists(d Directories, dn string, log *log.Logger) (*Directory, bool) {
	for _, dir := range d.ListDirect {
		if dir.Name == dn {
			log.Printf("found %s", dn)
			nd := &dir
			return nd, false
		}
	}
	log.Printf("Not found %s creating new directory", dn)
	nd := NewDirectory()
	nd.Name = dn
	return nd, true
}

func checkIfFileExists(d []File, dn string, log *log.Logger) (*File, bool) {
	for _, dir := range d {
		if dir.Name == dn {
			log.Printf("found %s", dn)
			return &dir, false
		}
	}
	log.Printf("Not found %s creating new File", dn)
	nf := NewFile()
	nf.Name = dn
	log.Println(nf.Name)
	return nf, true
}
