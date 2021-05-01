package db

import (
	"github.com/tannpv/bookstore_oauth-api/src/domain/access_token"
	"github.com/tannpv/bookstore_user-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implement yet")
}
