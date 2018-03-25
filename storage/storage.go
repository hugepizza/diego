package storage

import (
	"io"
	"os"
	"time"

	"github.com/ckeyer/diego/types"
)

type Storeger interface {
	GetUser(string) (*types.User, error)
	CreateUser(*types.User) error

	GetOrg(string) (*types.Org, error)
	CreateOrg(*types.Org) error
	// GetOrgs() (*types.Org, error)
}

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

// Namespace username or orgname
type Namespace struct {
	ID   string
	Name string
}

// Project
type Project struct {
	ID        string
	Namespace string
	Name      string
}

// Metadata
type FileMetadata struct {
	ID          string
	NamespaceID string
	ProjectID   string
	Name        string
	Size        int64
	Hash        string
	Mode        os.FileMode
	ModTime     time.Time
	CrdTime     time.Time
	IsDir       bool
	Version     string
	Labels      map[string]string
}
