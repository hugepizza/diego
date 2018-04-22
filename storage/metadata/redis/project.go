package redis

import (
	"time"

	"github.com/ckeyer/diego/types"
)

// CreateProject
func (r *RedisStorage) GetProject(name string) (*types.Project, error) {
	prj := &types.Project{}
	err := r.getKV(name, prj)
	if err != nil {
		return nil, err
	}
	return prj, nil
}

// CreateProject
func (r *RedisStorage) CreateProject(p *types.Project) error {
	p.Created = time.Now()
	// return r.setKV(p.Name, true)
	return nil
}

// CreateProject
func (r *RedisStorage) ListProjects(query string) ([]*types.Project, error) {
	ps := []*types.Project{}

	err := r.listKVs(query, &ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

// CreateProject
func (r *RedisStorage) DeleteProject() {

}
