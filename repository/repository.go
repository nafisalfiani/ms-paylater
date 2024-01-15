package repository

import (
	"ms-paylater/errors"

	"gorm.io/gorm"
)

type Repository struct {
	User UserInterface
}

func InitRepository(db *gorm.DB) *Repository {
	return &Repository{
		User: initUser(db),
	}
}

func errorAlias(err error) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.ErrNotFound
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return errors.ErrDuplicatedKey
	default:
		return err
	}
}
