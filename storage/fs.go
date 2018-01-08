package storage

import (
	"os"
)

type Metadatar interface {
	os.FileInfo
	List(string) ([]*Metadata, error)
}
