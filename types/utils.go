package types

const (
	OwnerTypeUser = "user"
	OwnerTypeOrg  = "org"
)

type Labels map[string]string

type ListOption struct {
	Offset int
	Limit  int
}

type ListUserOption struct {
	ListOption
}

type ListOrgOption struct {
	ListOption
}
