package writers

import (
	"os"
	"github.com/adagit94/gotils/fs"
)

func CreateByteFilesWriter(buffSize BuffSize, filePerms os.FileMode) *FilesWriter[[]byte, error] {
	fwPtr := &FilesWriter[[]byte, error]{PathsQueues: make(PathsQueues[error]), BuffSize: buffSize, Closed: make(chan bool, 1), SaveFile: func(path string, data []byte) error {
		return fs.Write(path, data, filePerms)
	}}

	return fwPtr
}
