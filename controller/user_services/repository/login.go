package userrepository

import (
	"database/sql"
	"fmt"

	usermodel "github.com/PhuPhuoc/koi-story-api-v2/controller/user_services/model"
)

func (store *userStore) Login(data *usermodel.UserLogin) (*usermodel.User, error) {
	user := &usermodel.User{}
	query := "select id, fb_id, email, password, name, avatar, role from users where email = ? and deleted_at is null"
	err := store.db.Get(user, query, data.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user does not exist")
		}
		return nil, fmt.Errorf("can not log in: %v", err)
	}

	if user.Password != "" && data.Password != user.Password {
		return nil, fmt.Errorf("wrong password")
	}
	return user, nil
}
