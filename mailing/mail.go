package mailing

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/s-vvardenfell/Backuper/utility"
)

type MailConfig struct {
	Sender   string   `json:"sender"`
	User     string   `json:"user"`
	Password string   `json:"passw"`
	Address  string   `json:"address"`
	Host     string   `json:"host"`
	To       []string `json:"to"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
	FileName string
}

type Mail struct {
}

// Send message with attachement with parameters, specified in config-file
// Recieves config file as "filename" arg
func (m *Mail) UploadFile(filename string) string {
	cnfg, err := LoadConfig("resources/TODO")

	if err != nil {
		log.Fatal(err)
	}

	SentMsgWithAttachment(cnfg)
	return ""
}

func LoadConfig(file string) (*MailConfig, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("cannot open config file: %w", err)
	}

	byteValue, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, fmt.Errorf("reading from file in LoadConfig failed: %w", err)
	}

	var res MailConfig
	json.Unmarshal([]byte(byteValue), &res)
	return &res, nil
}

func SentMsgWithAttachment(cnfg *MailConfig) {
	data := buildMail(cnfg)
	auth := smtp.PlainAuth("", cnfg.User, cnfg.Password, cnfg.Host)
	err := smtp.SendMail(cnfg.Address, auth, cnfg.Sender, cnfg.To, data)

	if err != nil {
		log.Fatal(err)
	}
}

func buildMail(mail *MailConfig) []byte {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", mail.Sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))

	boundary := "my-boundary-779"
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n",
		boundary))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", mail.Body))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString("Content-Disposition: attachment; filename=" + mail.FileName + "\r\n") //TODO убрать адрес/путь из названия файла с помощь ст библ
	buf.WriteString("Content-ID: <" + mail.FileName + ">\r\n\r\n")

	data := utility.ReadFile(mail.FileName)

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(b, data)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))
	buf.WriteString("--")
	return buf.Bytes()
}

func SendPlainMsg(cnfg *MailConfig) error {

	msg := []byte("From: " + cnfg.Sender + "\r\n" +
		"To: " + cnfg.To[0] + "\r\n" +
		"Subject: Test plain msg\r\n\r\n" +
		"Email body\r\n")

	auth := smtp.PlainAuth("", cnfg.User, cnfg.Password, cnfg.Host)

	if err := smtp.SendMail(cnfg.Address, auth, cnfg.Sender, cnfg.To, msg); err != nil {
		return err
	}
	return nil
}
