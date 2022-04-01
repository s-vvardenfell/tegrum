package storages

import "io"

// TODO интерфейс туда, где он исп-ся
type Storage interface {
	Store(w io.Writer, data []string) error
	Retrieve(r io.Reader, index string) ([]string, error)
}
