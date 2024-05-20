package repository

import (
	"database/sql"

	"github.com/mass584/twitter-clone-app/backend/domain/entity"
	"github.com/mass584/twitter-clone-app/backend/domain/repository_if"
	"github.com/go-sql-driver/mysql"
)

type FollowableUserRepository struct {
	database *sql.DB
}

func NewFollowableUserRepository(db *sql.DB) repository_if.FollowedUserRepository {
	return &FollowableUserRepository{database: db}
}

func (r *FollowableUserRepository) Create(user_id int64, target_user_id int64) error {
	statement, _ := r.database.Prepare(
		"insert into follows (follower_user_id, followed_user_id) values (?, ?)",
	)
	defer statement.Close()

	_, err := statement.Exec(user_id, target_user_id)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return nil
			}
		}
		return err
	}

	return nil
}

func (r *FollowableUserRepository) Delete(user_id int64, target_user_id int64) error {
	statement, _ := r.database.Prepare(
		"delete from follows where follower_user_id = ? and followed_user_id = ?",
	)
	defer statement.Close()

	_, err := statement.Exec(user_id, target_user_id)
	if err != nil {
		return err
	}

	return nil
}

func (r *FollowableUserRepository) FindListByQuery(user_id int64, query string) (*[]entity.FollowableUser, error) {
	statement, _ := r.database.Prepare(
		"select users.id, users.display_name, follows.followed_user_id from users " +
			"left join (select follower_user_id, followed_user_id from follows where follower_user_id = ?) as follows " +
			"on follows.followed_user_id = users.id " +
			"where users.display_name like ? and users.id <> ? " +
			"order by users.created_at desc ",
	)
	defer statement.Close()

	pattern := "%" + query + "%"
	rows, _ := statement.Query(user_id, pattern, user_id)
	defer rows.Close()

	var followable_users []entity.FollowableUser = []entity.FollowableUser{}

	for rows.Next() {
		var id int64
		var display_name string
		var followed_user_id int64 = 0
		var is_followed bool

		rows.Scan(&id, &display_name, &followed_user_id)
		user := entity.User{Id: id, DisplayName: display_name}
		if followed_user_id == 0 {
			is_followed = false
		} else {
			is_followed = true
		}

		followable_user, _ := entity.NewFollowableUser(user, is_followed)
		followable_users = append(followable_users, *followable_user)
	}

	return &followable_users, nil
}
