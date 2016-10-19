package store

import "github.com/praelatus/backend/models"

type sqlProjectStore struct{}

func (sp *sqlProjectStore) Get(id string) models.Project {
	return models.Project{}
}
