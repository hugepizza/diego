package storage

import (
	"os"
	"time"
)

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
