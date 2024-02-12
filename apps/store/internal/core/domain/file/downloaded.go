package file

type Downloaded struct {
	DownloadPath string
	Meta         File
}

func NewDownloaded(path string, meta File) Downloaded {
	return Downloaded{path, meta}
}
