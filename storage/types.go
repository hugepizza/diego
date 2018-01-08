package storage

import (
	"os"
	"time"

	"github.com/Masterminds/semver"
)

type FileInfo struct {
	Name    string
	Size    int64
	Version string
	Hash    string
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

type VersionDir struct {
	Namespace string
	Version   *semver.Version
}
