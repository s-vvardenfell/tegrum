package taring

import (
	"testing"

	"github.com/s-vvardenfell/tegrum/utility"
	"github.com/stretchr/testify/require"
)

const isNestedPkg = true

func TestTarUnTar(t *testing.T) {
	tar := Tar{}
	var archiveLocation string
	tempDir, _, tempFileName, err := utility.PrepareForTest(isNestedPkg)
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
