package commentrepository

import (
	"fmt"

	"github.com/PhuPhuoc/koi-story-api-v2/utils"
)

func (store *commentStore) DeleteComment(comment_id string) error {
	date := utils.GetCurrentDateTime()
	query := `
		update comments set deleted_at=? where id=?
	`
	if _, err := store.db.Exec(query, date, comment_id); err != nil {
		return fmt.Errorf("error when delete comment: %w", err)
	}
	return nil
}
