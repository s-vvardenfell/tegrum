package cmd

import "io"

type Extensioner interface {
	Extension() string
}

type Uploader interface {
	Extensioner
	UploadFile(filename string) (string, error)
}

type Downloader interface {
	Extensioner
	DownLoadFile(fileId, dst string) error
}

type Archiver interface {
	Extensioner
	Archive(source, dst string) (string, error)
}

type Extracter interface {
	Extensioner
	Extract(archive, dst string) error
}

type RecorderRetriever interface {
	Record(w io.Writer, data []string) error
	Retrieve(r io.Reader, index string) ([]string, error)
}

type DirsToBackup struct {
	Dirs []string `json:"dirs"`
}
