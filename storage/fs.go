package storage

import (
	"io"
	"os"
)

type FSDriver interface {
	DataStoreger
	MetadataStoreger
}

type DataStoreger interface {
	Write(info *FileMetadata, data io.Reader) error
}

type MetadataStoreger interface {
	os.FileInfo
	List(string) ([]*FileMetadata, error) // ls
}

func CreateFile(fs FSDriver, info *FileMetadata, data io.Reader) {
	return
}
