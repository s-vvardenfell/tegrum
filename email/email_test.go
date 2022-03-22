package email

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const resourceDir = "W:/Golang/src/Backuper/resources"
const configName = "email.json"
const attachFile = "test_file_48kb.txt"
const subject = "testing email-pkg"
const body = "test body msg"

func TestUploadFile(t *testing.T) {
	e := NewMail(filepath.Join(resourceDir, configName))
	t.Log("\tSending plain message with attachment")
	{
		_, err := e.UploadFile(filepath.Join(resourceDir, attachFile))
		require.NoError(t, err)
	}
}

func TestSendPlainMsg(t *testing.T) {
	e := NewMail(filepath.Join(resourceDir, configName))
	t.Log("\tSending plain message with attachment")
	{
		err := e.SendPlainMsg(subject, body)
		require.NoError(t, err)
	}
}
