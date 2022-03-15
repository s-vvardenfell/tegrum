package archiver

import "fmt"

type Tar struct {
}

func (t *Tar) Archive(source, target string) error {
	fmt.Println("Works TAR")
	return nil
}

func (t *Tar) Extract(archive, target string) error {
	fmt.Println("Works TAR")
	return nil
}
