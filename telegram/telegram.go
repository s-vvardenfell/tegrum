package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

const ext = "telegram"

var baseUrl = "https://api.telegram.org/bot"
var fileUrl = "https://api.telegram.org/file/bot"
var timeout = 30

func NewTelegram(config string) *Telegram {
	filename, err := filepath.Abs(config)
	if err != nil {
		logrus.Fatal(err)
	}

	jFile, err := os.ReadFile(filename)
	if err != nil {
		logrus.Fatal(err)
	}

	var tg Telegram
	if err := json.Unmarshal(jFile, &tg); err != nil {
		logrus.Fatal(err)
	}
	tg.extension = ext
	return &tg
}

func (t *Telegram) Extension() string {
	return t.extension
}

// Downloads a file from a telegram chat using a chat-bot
// specified in the configuration
// NOTE: files larger than 20 MB may not be downloaded because of API restrictions
func (t *Telegram) DownLoadFile(fileId, dst string) error {
	url, err := fileLocationFromServer(t.Token, fileId)
	if err != nil {
		return err
	}

	if err := downloadFileFromServer(url, dst); err != nil {
		return err
	}
	return nil
}

// Uploads files up to 50 MB to the chat
// using a chat bot specified in the configuration
// returns file id from telegram server
func (t *Telegram) UploadFile(filename string) (string, error) {
	if err := checkSize(filename); err != nil {
		return "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fw, _ := writer.CreateFormFile("document", filename)
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("cannot open %s, %w", filename, err)
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return "", err
	}

	writer.Close()

	url := fmt.Sprintf("%s%s/sendDocument?chat_id=%s", baseUrl, t.Token, t.ChatId)

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tr TelegramUploadResponse
	if err := json.Unmarshal(byteValue, &tr); err != nil {
		return "", err
	}

	if tr.Ok {
		return tr.Result.Document.FileId, nil
	}
	return "", errors.New("file uploading to telegram error")
}

func checkSize(filename string) error {
	files, err := ioutil.ReadDir(filepath.Dir(filename))
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == filepath.Base(filename) {
			if file.Size() > 52428800 {
				return errors.New("cannot upload files larger that 50 MB in Telegram")
			}
		}
	}
	return nil
}

func fileLocationFromServer(token, fileId string) (string, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	url := fmt.Sprintf("%s%s/getFile?file_id=%s", baseUrl, token, fileId)
	bytesValue, err := doRequest(client, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("doRequest() failed, %v", err)
	}

	var tr TelegramDownloadResponse
	if err := json.Unmarshal(bytesValue, &tr); err != nil {
		return "", err
	}

	if tr.Ok {
		return fmt.Sprintf("%s%s/%s", fileUrl, token, tr.Result.FilePath), nil
	}
	return "", errors.New("cannot get url with file location on telegram server")
}

func downloadFileFromServer(url, dst string) error {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	bytesValue, err := doRequest(client, http.MethodGet, url, nil)

	if err != nil {
		return fmt.Errorf("doRequest() failed, %v", err)
	}

	//trying to parse response; if can, there is an error
	var de TelegramDownloadError
	if err := json.Unmarshal(bytesValue, &de); err == nil {
		return fmt.Errorf("file download error; code:%v, descr:%s; %w", de.ErrorCode, de.Descr, err)
	}

	if err := ioutil.WriteFile(filepath.Join(dst, filepath.Base(url)), bytesValue, 0644); err != nil {
		return fmt.Errorf("cannot write result file, %v", err)
	}
	return nil
}

func doRequest(client *http.Client, method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bytesValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytesValue, nil
}
