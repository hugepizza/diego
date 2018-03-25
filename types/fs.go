package types

import (
	"fmt"
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
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Desc  string `json:"desc"`

	Created time.Time `json:"created"`
}

// Metadata
type FileMetadata struct {
	Name    string            `json:"name"`
	Owner   string            `json:"owner"`
	Project string            `json:"project"`
	Path    string            `json:"path"`
	Size    int64             `json:"size"`
	Hash    string            `json:"hash"`
	Type    string            `json:"type"`
	Mode    os.FileMode       `json:"mode"`
	ModTime time.Time         `json:"modified"`
	CrdTime time.Time         `json:"created"`
	Version string            `json:"version"`
	Status  string            `json:"status"`
	Labels  map[string]string `json:"labels"`
}

type FilePath struct {
	Hash string
	Path string
}

func (p *Project) Prefix() string {
	return "prj:"
}

// Key
func (p *Project) Key() string {
	return fmt.Sprintf("%s:%s:%s", p.Prefix(), p.Owner, p.Name)
}

// Key
func (f *FileMetadata) Prefix() string {
	return "fileinfo:"
}

// Key
func (f *FileMetadata) Key() string {
	path := f.Name
	if f.Path != "" {
		path = f.Path
	}
	return fmt.Sprintf("%s:%s:%s:%s@%s", f.Prefix(), f.Owner, f.Project, path, f.Version)
}

// Key
func (f *FilePath) Prefix() string {
	return "filepath:"
}

// Key
func (f *FilePath) Key() string {
	return f.Prefix() + f.Hash
}
