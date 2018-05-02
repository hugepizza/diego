package types

import (
	"os"
	"time"
)

const (
	FileStatusOK        = "ok"
	FileStatusUploading = "uploading"
	FileStatusDeleted   = "deleted"
)
const (
	FileTypeBin      = "binary"
	FileTypeMarkdown = "markdown"
	FileTypeText     = "text"
)

// Project
type Project struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Desc      string `json:"desc"`
	Labels    Labels `json:"labels"`

	Created time.Time `json:"created"`
}

// Metadata
type FileMetadata struct {
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	ProjectName string      `json:"project_name"`
	Size        int64       `json:"size"`
	Hash        string      `json:"hash"`
	Type        string      `json:"type"`
	Mode        os.FileMode `json:"mode"`
	ModTime     time.Time   `json:"modified"`
	CrdTime     time.Time   `json:"created"`
	Version     string      `json:"version"`
	Status      string      `json:"status"`
	Labels      Labels      `json:"labels"`
}
