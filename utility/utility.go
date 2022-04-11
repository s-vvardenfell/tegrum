package utility

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func ReadFile(fileName string) []byte {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return bytes
}

func WriteFile(fileName, content string) {
	err := os.WriteFile(fileName, []byte(content), 0666)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func PrepareForTest(nestedPkg bool) (tempDir string, resourceDir string, tempFileName string, err error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", "", "", err
	}

	//get temp/resource dir
	if nestedPkg {
		tempDir = filepath.Join(filepath.Join(filepath.Dir(wd), "../"), "temp")
		resourceDir = filepath.Join(filepath.Join(filepath.Dir(wd), "../"), "resources")
	} else {
		tempDir = filepath.Join(filepath.Dir(wd), "temp")
		resourceDir = filepath.Join(filepath.Dir(wd), "resources")
	}

	tempFileName = filepath.Join(tempDir, time.Now().Format("02-01-2006_15-04-05")+".txt") //create temp file to upload
	tempFileContent := strconv.Itoa(int(time.Now().Unix()))                                //fill temp file with content
	if err = os.WriteFile(tempFileName, []byte(tempFileContent), 0666); err != nil {
		return "", "", "", err
	}
	return tempDir, resourceDir, tempFileName, nil
}
