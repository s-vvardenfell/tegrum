package telegram

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const resourceDir = "W:/Golang/src/Backuper/resources"
const resultDir = "W:/Golang/src/Backuper/result"
const configName = "telegram.json"
const testFile = "test_file.txt"

func TestUploadDownload(t *testing.T) {

	tg := NewTelegram(filepath.Join(resourceDir, configName))
	var fileId string
	var err error

	t.Log("\tUploading file to telegramm chat")
	{
		fileId, err = tg.UploadFile(filepath.Join(resourceDir, testFile))
		require.NoError(t, err)
		require.True(t, (fileId != ""))
	}

	t.Log("\tDownloading file from telegram chat")
	{
		url, err := fileLocationFromServer(tg.Token, fileId)
		require.True(t, (url != ""))
		require.NoError(t, err)

		err = downloadFileFromServer(url, resultDir)
		require.NoError(t, err)
		require.FileExists(t, filepath.Join(resultDir, filepath.Base(url)))
	}
}
