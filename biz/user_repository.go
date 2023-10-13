package biz

import (
	"database/sql"
	"echo/biz/domain"
)

type User struct {
	Id             int64
	Name           string
	HashedPassword string
	ApiToken       string
}

type UserRepositoryContract interface {
	Insert(user domain.User) error
	GetByName(name string) (domain.User, error)
}


type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Insert(user domain.User) error {
	query := "INSERT INTO echotraceschema.user (name, hashed_password, api_token) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, user.Name, user.HashedPassword, user.ApiToken)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetByName(name string) (domain.User, error) {
	query := "SELECT id, name, hashed_password, api_token FROM echotraceschema.user WHERE name = ?"
	row := r.db.QueryRow(query, name)

	var user domain.User
	err := row.Scan(&user.Id, &user.Name, &user.HashedPassword, &user.ApiToken)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

