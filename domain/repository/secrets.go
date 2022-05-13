package repository

import "github.com/haradayoshitsugucz/purple-server/domain/entity"

type SecretsRepository interface {
	GetDBCredentials(secretName string) (*entity.DBCredentials, error)
}
