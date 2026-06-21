package writers

import (
	"os"
	"github.com/adagit94/gotils/fs"
)

func CreateByteFilesWriter(buffSize buffSize, filePerms os.FileMode) IFilesWriter[[]byte, error] {
	fwPtr := &filesWriter[[]byte, error]{pathsQueues: make(pathsQueues[error]), buffSize: buffSize, saveFile: func(path string, data []byte) error {
		return fs.Write(path, data, filePerms)
	}}

	return fwPtr
}
