package utility

import (
	"fmt"
	"os"
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
