package repository

import (
	"github.com/haradayoshitsugucz/purple-server/domain/model"
)

type SecretsRepository interface {
	GetDBCredentials(secretName string) (*model.DBCredentials, error)
}
