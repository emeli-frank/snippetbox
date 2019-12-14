package mysql

import (
	"database/sql"
	"emeli/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

func (m *UserModel) Authenticate(name, email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(name, email, password string) (*models.User, error) {
	return nil, nil
}
