package main

import "github.com/s-vvardenfell/Backuper/cmd"

type DirsToBackup struct {
	Dirs []string `json:"dirs"`
}

func main() {
	cmd.Execute()

	// f, _ := os.Open("resources/map.json")

	// byteValue, _ := ioutil.ReadAll(f)

	// var dirs DirsToBackup

	// json.Unmarshal([]byte(byteValue), &dirs)

	// for _, dir := range dirs.Dirs {
	// 	archiver.Zip(dir, "result/"+filepath.Base(dir)+".zip")
	// }
}
