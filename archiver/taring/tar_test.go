package archiver

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTarUnTar(t *testing.T) {
	wd, err := os.Getwd()
	require.NoError(t, err)
	tempDir := filepath.Join(filepath.Join(filepath.Dir(wd), "../"), "temp")

	tar := Tar{}
	tempFileName := filepath.Join(tempDir, time.Now().Format("02-01-2006_15-04-05")+".txt")
	tempFileContent := strconv.Itoa(int(time.Now().Unix()))
	var archiveLocation string

	err = os.WriteFile(tempFileName, []byte(tempFileContent), 0666)
	require.NoError(t, err)

	t.Log("\tTaring/gzipping file")
	{
		archiveLocation, err = tar.Archive(tempFileName, tempDir)
		require.NoError(t, err)
		require.NotEmpty(t, archiveLocation)
		require.FileExists(t, archiveLocation)
	}

	t.Log("\tUntaring/gzippong file")
	{
		err = tar.Extract(archiveLocation, tempDir)
		require.NoError(t, err)
		require.NotEmpty(t, archiveLocation)
		require.FileExists(t, archiveLocation)
	}
}
