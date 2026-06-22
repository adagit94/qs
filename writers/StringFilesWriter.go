package writers

import (
	"os"
	"github.com/adagit94/gotils/fs"
)

func CreateStringFilesWriter(buffSize BuffSize, perms os.FileMode) *FilesWriter[string, error] {
	fwPtr := &FilesWriter[string, error]{PathsQueues: make(PathsQueues[error]), BuffSize: buffSize, Closed: make(chan bool, 1), SaveFile: func(path string, data string) error {
		return fs.WriteString(path, data, perms)
	}}

	return fwPtr
}
