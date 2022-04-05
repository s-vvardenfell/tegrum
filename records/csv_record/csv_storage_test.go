package csv_record

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testFile = "test_data.csv"

func TestStoreRetrieve(t *testing.T) {
	st := CsvStorage{}
	test_data := strconv.FormatInt(time.Now().Unix(), 10)   //generate data for slice
	test_slice := []string{test_data, test_data, test_data} //data to write

	wd, err := os.Getwd()
	require.NoError(t, err)
	tempDir := filepath.Join(filepath.Join(filepath.Dir(wd), "../"), "temp") //get temp dir

	t.Log("\tStore record to .csv file")
	{
		file, err := os.OpenFile(filepath.Join(tempDir, testFile), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		require.NoError(t, err)
		defer func() { _ = file.Close }()

		err = st.Store(file, test_slice)
		require.NoError(t, err)
	}

	t.Log("\tRead record from .csv file")
	{
		file, err := os.OpenFile(filepath.Join(tempDir, testFile), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		require.NoError(t, err)
		defer func() { _ = file.Close }()

		got, err := st.Retrieve(file, test_data)
		require.NoError(t, err)
		require.Contains(t, got, test_data)
	}
}
