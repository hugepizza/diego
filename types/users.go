package types

import (
	"time"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`

	Joined time.Time `json:"joined"`
}

func (u *User) Prefix() string {
	return "user:"
}

// Key
func (u *User) Key() string {
	return u.Prefix() + u.Name
}

type Org struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`

	Created time.Time `json:"created"`
}

func (o *Org) Prefix() string {
	return "org:"
}

// Key
func (o *Org) Key() string {
	return o.Prefix() + o.Name
}
