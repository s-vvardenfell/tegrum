package telegram

import "fmt"

type Telegram struct {
}

func (t *Telegram) DownLoadFile(fileId, dst string) {
	fmt.Println("Downloading file from telegram")
}

func (t *Telegram) UploadFile(filename string) string {
	fmt.Println("Uploading file to telegram")
	return ""
}
