package email

import (
	"path/filepath"
	"testing"

	"github.com/s-vvardenfell/tegrum/utility"
	"github.com/stretchr/testify/require"
)

const configName = "email.json"
const subject = "testing email-pkg"
const body = "test body msg"
const isNestedPkg = false

func TestUploadFile(t *testing.T) {
	tempDir, resourceDir, tempFileName, err := utility.PrepareForTest(isNestedPkg)
	require.NoError(t, err)

	e := NewMail(filepath.Join(resourceDir, configName))
	t.Log("\tSending message with attachment")
	{
		_, err := e.UploadFile(filepath.Join(tempDir, tempFileName))
		require.NoError(t, err)
	}
}

func TestSendPlainMsg(t *testing.T) {
	_, resourceDir, _, err := utility.PrepareForTest(isNestedPkg)
	require.NoError(t, err)

	e := NewMail(filepath.Join(resourceDir, configName))
	t.Log("\tSending plain message")
	{
		err := e.SendPlainMsg(subject, body)
		require.NoError(t, err)
	}
}
