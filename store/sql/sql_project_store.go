package store

import "github.com/chasinglogic/tessera/models"

type sqlProjectStore struct{}

func (sp *sqlProjectStore) Get(id string) models.Project {
	return models.Project{}
}
