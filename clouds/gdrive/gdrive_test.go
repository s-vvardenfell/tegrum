package gdrive

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const credentials = ""
const fileToUpload = "" //должен генерироваться
const dstDir = ""

//TODO
// const isNestedPkg = true
// tempDir, _, tempFileName, err := utility.PrepareForTest(isNestedPkg)
// require.NoError(t, err)

func TestUploadDownload(t *testing.T) {
	gd := NewGDrive(credentials)
	var fileId string
	var fileName string
	var err error

	t.Log("\tUpload archive to Google Drive")
	{
		fileId, err = gd.UploadFile(fileToUpload)
		require.NoError(t, err)
		require.NotEmpty(t, fileId)
	}

	t.Log("\tGetting archive name by stored file id from Google Drive")
	{
		fileName, err = gd.fileNameById(fileId)
		require.NoError(t, err)
		require.NotEmpty(t, fileName)
	}

	t.Log("\tDownload archive from Google Drive")
	{
		err = gd.DownLoadFile(fileId, dstDir)
		require.NoError(t, err)
		require.FileExists(t, filepath.Join(dstDir, fileName))
	}

	t.Log("\tGetting archive file id by file name from Google Drive")
	{
		tempName, err := gd.fileIdByName(fileName)
		require.NoError(t, err)
		require.EqualValues(t, fileName, tempName)
	}

	t.Log("\tDeleting temporary file by file id from Google Drive")
	{
		err = gd.deleteFile(fileId)
		require.NoError(t, err)
	}

	t.Log("\tChecking file was really removed from Google Drive")
	{
		err = gd.DownLoadFile(fileId, dstDir)
		require.Error(t, err)
	}
}
