package zipping

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestZipUnzip(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	tempDir := filepath.Join(filepath.Join(filepath.Dir(wd), "../"), "temp")

	zip := Zip{}
	tempFileName := filepath.Join(tempDir, time.Now().Format("02-01-2006_15-04-05")+".txt")
	tempFileContent := strconv.Itoa(int(time.Now().Unix()))
	var archiveLocation string

	err = os.WriteFile(tempFileName, []byte(tempFileContent), 0666)
	require.NoError(t, err)

	t.Log("\tZipping file")
	{
		archiveLocation, err = zip.Archive(tempFileName, tempDir)
		require.NoError(t, err)
		require.NotEmpty(t, archiveLocation)
		require.FileExists(t, archiveLocation)
	}

	t.Log("\tUnzipping file")
	{
		err = zip.Extract(archiveLocation, tempDir)
		require.NoError(t, err)
		require.NotEmpty(t, archiveLocation)
		require.FileExists(t, archiveLocation)
	}

	// error - "open in another process"
	// t.Log("\tCleaning")
	// {
	// 	err = os.Remove(archiveLocation)
	// 	require.NoError(t, err)

	// 	err = os.Remove(tempFileName)
	// 	require.NoError(t, err)
	// }
}
