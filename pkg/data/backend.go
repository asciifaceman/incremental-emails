package data

import (
	"io/fs"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

const DEFAULT_TIMEOUT = 1
const DEFAULT_TIMEOUT_FORMAT = time.Second
const DEFAULT_MODE = 0666

func NewBackend(path string, mode fs.FileMode) *Backend {
	return &Backend{
		Path:    path,
		Mode:    mode,
		Options: &bolt.Options{Timeout: DEFAULT_TIMEOUT * DEFAULT_TIMEOUT_FORMAT},
	}
}

// Backend is the
type Backend struct {
	Data    *bolt.DB
	Path    string
	Mode    fs.FileMode
	Options *bolt.Options
}

func (b *Backend) NoSave() bool {
	_, err := os.Stat(b.Path)
	return err != nil
}
