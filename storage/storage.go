package storage

import (
	"github.com/ckeyer/diego/storage/metadata"
)

type Storager interface {
	metadata.UserStorager
	metadata.OrgStorager
	metadata.NamespaceStorager
	metadata.ProjectStorager
	metadata.FileIndexer
}

type Keyer interface {
	Prefix() string
	Key() string
}
