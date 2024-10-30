package commentrepository

import (
	"fmt"

	commentmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/model"
)

func (store *commentStore) GetAllComment(post_id string) (data []commentmodel.Comment, err error) {
	data = []commentmodel.Comment{}
	tx, err := store.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("cannot begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // Rethrow panic after rollback.
		} else if err != nil {
			_ = tx.Rollback() // Ignore rollback error and return original error.
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction: %w", commitErr)
		}
	}()

	query_comment := `
		select c.id, c.user_id, u.name, u.avatar, c.content, c.created_at
		from comments c
		join users u on c.user_id=u.id
		where post_id=? and c.deleted_at is null
	`
	if err = tx.Select(&data, query_comment, post_id); err != nil {
		return nil, fmt.Errorf("cannot get comments: %w", err)
	}
	return data, nil
}
