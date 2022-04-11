package gdrive

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/s-vvardenfell/tegrum/utility"
	"github.com/stretchr/testify/require"
)

const credentials = "credentials.json"
const isNestedPkg = true

// gdrive api saves token.json with access and refresh tokens to the root dir
// in tests it is not possible to follow link in browser, press buttons to allow access to and get this token
// so, this test must be run manually and won't be included in CI
func TestUploadDownload(t *testing.T) {

	tempDir, resourceDir, tempFileName, err := utility.PrepareForTest(isNestedPkg)
	fmt.Println(tempDir, resourceDir, tempFileName)
	require.NoError(t, err)
	gd := NewGDrive(filepath.Join(resourceDir, credentials))
	var fileId string
	var fileName string

	t.Log("\tUpload archive to Google Drive")
	{
		fileId, err = gd.UploadFile(tempFileName)
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
		err = gd.DownLoadFile(fileName, tempDir)
		require.NoError(t, err)
		require.FileExists(t, filepath.Join(tempDir, fileName))
	}

	t.Log("\tGetting archive file id by file name from Google Drive")
	{
		tempName, err := gd.fileIdByName(fileName)
		require.NoError(t, err)
		require.EqualValues(t, fileId, tempName)
	}

	t.Log("\tDeleting temporary file by file id from Google Drive")
	{
		err = gd.deleteFile(fileId)
		require.NoError(t, err)
	}

	t.Log("\tChecking file was really removed from Google Drive")
	{
		err = gd.DownLoadFile(fileId, tempDir)
		require.Error(t, err)
	}
}
