package file

import "os"

type Downloaded struct {
	File *os.File
	Meta File
}

func NewDownloaded(meta File, file *os.File) Downloaded {
	return Downloaded{file, meta}
}
