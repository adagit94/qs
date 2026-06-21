package writers

import (
	"os"

	"github.com/adagit94/gotils/fs"
)

func CreateStringFilesWriter(buffSize buffSize, perms os.FileMode) IFilesWriter[string, error] {
	fwPtr := &filesWriter[string, error]{pathsQueues: make(pathsQueues[error]), buffSize: buffSize, saveFile: func(path string, data string) error {
		return fs.WriteString(path, data, perms)
	}}

	return fwPtr
}
