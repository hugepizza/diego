package storage

import (
	"github.com/ckeyer/diego/types"
)

type Storeger interface {
	GetUser(string) (*types.User, error)
	CreateUser(*types.User) error
	ListUsers() ([]*types.User, error)

	GetOrg(string) (*types.Org, error)
	CreateOrg(*types.Org) error
	ListOrgs() ([]*types.Org, error)

	GetProject(string) (*types.Project, error)
	CreateProject(*types.Project) error
}

type Keyer interface {
	Prefix() string
	Key() string
}
