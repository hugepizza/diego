package storage

import (
	"github.com/ckeyer/diego/types"
)

type Storeger interface {
	ExistsUser(string) (bool, error)
	GetUser(string) (*types.User, error)
	CreateUser(*types.User) error
	ListUsers() ([]*types.User, error)

	GetOrg(string) (*types.Org, error)
	CreateOrg(*types.Org) error
	ListOrgs() ([]*types.Org, error)

	// ExistsProject(string)

	GetProject(string) (*types.Project, error)
	ListProjects(string) ([]*types.Project, error)
	CreateProject(*types.Project) error
}

type Keyer interface {
	Prefix() string
	Key() string
}
