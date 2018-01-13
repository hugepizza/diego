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
	Write(info *Metadata, data io.Reader) error
}

type MetadataStoreger interface {
	os.FileInfo
	List(string) ([]*Metadata, error) // ls
}

func CreateFile(fs FSDriver, info *Metadata, data io.Reader) {

}
