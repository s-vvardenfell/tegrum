package yadisk

import "fmt"

const ext = "yadisk"

type YaDisk struct {
	extension string
}

func NewYaDisk(config string) *YaDisk {
	return &YaDisk{ext}
}

func (yd *YaDisk) Extension() string {
	return yd.extension
}

func (yd *YaDisk) DownLoadFile(fileId, dst string) error {
	fmt.Println("Downloading archive from Yandex Disk")
	return fmt.Errorf("Yandex-Disk not implemented")
}

func (yd *YaDisk) UploadFile(filename string) (string, error) {
	fmt.Println("Uploading archive to Yandex Disk")
	return "", fmt.Errorf("Yandex-Disk not implemented")
}
