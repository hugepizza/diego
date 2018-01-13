package storage

import (
	"os"
	"time"

	"github.com/Masterminds/semver"
)

type Metadata struct {
	Name    string
	Size    int64
	Hash    string
	Mode    os.FileMode
	ModTime time.Time
	CrdTime time.Time
	IsDir   bool
	Version *semver.Version
}
