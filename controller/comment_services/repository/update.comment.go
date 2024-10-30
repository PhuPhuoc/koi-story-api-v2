package commentrepository

import (
	"fmt"

	commentmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/model"
)

func (store *commentStore) UpdateComment(comment_id string, data *commentmodel.UpdateComment) error {
	data.ID = comment_id
	query := `
		update comments set content=:content where id=:id
	`
	if _, err := store.db.NamedExec(query, data); err != nil {
		return fmt.Errorf("error when update comment: %w", err)
	}
	return nil
}
