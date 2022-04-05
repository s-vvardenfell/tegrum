package zipping

import (
	"testing"

	"github.com/s-vvardenfell/tegrum/utility"
	"github.com/stretchr/testify/require"
)

const isNestedPkg = true

func TestZipUnzip(t *testing.T) {
	zip := Zip{}
	var archiveLocation string
	tempDir, _, tempFileName, err := utility.PrepareForTest(isNestedPkg)
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
