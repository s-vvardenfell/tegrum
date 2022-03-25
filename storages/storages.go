package storages

import "io"

type Storage interface {
	Store(w io.Writer, data []string) error
	Retrieve(r io.Reader, index string) ([]string, error)
}
