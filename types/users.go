package types

import (
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`

	Joined  time.Time `json:"joined"`
	Updated time.Time `json:"updated"`
}

type Org struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`
	Owner string `json:"owner"`

	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type Namespace struct {
	Name      string `json:"name"`
	OwnerType string `json:"owner_type"`
	Labels    Labels `json:"labels"`
}
