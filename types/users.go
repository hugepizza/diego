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

type Org struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Desc  string `json:"desc"`

	Created time.Time `json:"joined"`
}
