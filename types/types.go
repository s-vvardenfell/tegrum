package types

type Uploader interface {
	UploadFile(filename string) (string, error)
}

type Downloader interface {
	DownLoadFile(fileId, dst string) error
}
