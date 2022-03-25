package storages

import "io"

type Storage interface {
	Store(r *io.Reader) error
	Retrieve(r io.Reader, index string) ([]string, error)
}
