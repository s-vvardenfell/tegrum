package email

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"

	"github.com/s-vvardenfell/Backuper/utility"
	"github.com/sirupsen/logrus"
)

const ext = "email"

type Email struct {
	Sender    string   `json:"sender"`
	User      string   `json:"user"`
	Password  string   `json:"passw"`
	Address   string   `json:"address"`
	Host      string   `json:"host"`
	To        []string `json:"to"`
	Subject   string   `json:"subject"`
	Body      string   `json:"body"`
	extension string
}

func NewMail(cnfg string) *Email {
	f, err := os.Open(cnfg)
	if err != nil {
		logrus.Fatalf("cannot open config file: %v", err)
	}

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		logrus.Fatalf("reading from file in LoadConfig failed: %v", err)
	}

	var res Email
	json.Unmarshal([]byte(byteValue), &res)
	res.extension = ext
	return &res
}

func (m *Email) Extension() string {
	return m.extension
}

// Sends message with attachement with parameters, specified in config-file
func (m *Email) UploadFile(filename string) (string, error) {
	data := m.buildMail(filename)
	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)
	err := smtp.SendMail(m.Address, auth, m.Sender, m.To, data)
	return ext, err
}

func (m *Email) buildMail(filename string) []byte {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", m.Sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(m.To, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", m.Subject))

	boundary := "my-boundary-779"
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n",
		boundary))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", m.Body))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString("Content-Disposition: attachment; filename=" + filepath.Base(filename) + "\r\n")
	buf.WriteString("Content-ID: <" + filepath.Base(filename) + ">\r\n\r\n")

	data := utility.ReadFile(filename)

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(b, data)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))
	buf.WriteString("--")
	return buf.Bytes()
}

func (m *Email) SendPlainMsg(subject, body string) error {
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		m.Sender, m.To[0], subject, body))

	auth := smtp.PlainAuth("", m.User, m.Password, m.Host)
	err := smtp.SendMail(m.Address, auth, m.Sender, m.To, msg)
	return err
}
