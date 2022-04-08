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
	return fmt.Errorf("Yandex-Disk not implemented")
}

func (yd *YaDisk) UploadFile(filename string) (string, error) {
	return "", fmt.Errorf("Yandex-Disk not implemented")
}
