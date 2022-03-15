package clouds

import "fmt"

type YaDisk struct {
}

func NewYaDisk() *YaDisk {
	return &YaDisk{}
}

func (yd *YaDisk) DownLoadFile(fileId, dst string) {
	fmt.Println("Downloading archive from Yandex Disk")
}

func (yd *YaDisk) UploadFile(filename string) string {
	fmt.Println("Uploading archive to Yandex Disk")
	return ""
}
