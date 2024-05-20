package repository

import (
	"database/sql"
	"fmt"

	"github.com/mass584/twitter-clone-app/backend/domain/entity"
	"github.com/mass584/twitter-clone-app/backend/domain/repository_if"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(db *sql.DB) repository_if.UserRepository {
	return &UserRepository{database: db}
}

func (r *UserRepository) FindOneByEmail(email string) (*entity.User, error) {
	statement, _ := r.database.Prepare("select id, display_name from users where email = ?")
	defer statement.Close()
	rows, err := statement.Query(email)
	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, fmt.Errorf("User is not found.")
	}

	var id int64
	var display_name string
	rows.Scan(&id, &display_name)

	user, _ := entity.NewUser(id, display_name)

	return user, nil
}
