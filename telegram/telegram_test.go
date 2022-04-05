package telegram

import (
	"path/filepath"
	"testing"

	"github.com/s-vvardenfell/tegrum/utility"
	"github.com/stretchr/testify/require"
)

const configName = "telegram.json"
const isNestedPkg = false

func TestUploadDownload(t *testing.T) {
	tempDir, resourceDir, tempFileName, err := utility.PrepareForTest(isNestedPkg)
	require.NoError(t, err)

	tg := NewTelegram(filepath.Join(resourceDir, configName))
	var fileId string

	t.Log("\tUploading file to telegramm chat")
	{
		fileId, err = tg.UploadFile(tempFileName)
		require.NoError(t, err)
		require.True(t, (fileId != ""))
	}

	t.Log("\tDownloading file from telegram chat")
	{
		url, err := fileLocationFromServer(tg.Token, fileId)
		require.True(t, (url != ""))
		require.NoError(t, err)

		err = downloadFileFromServer(url, tempDir)
		require.NoError(t, err)
		require.FileExists(t, filepath.Join(tempDir, filepath.Base(url)))
	}
}
