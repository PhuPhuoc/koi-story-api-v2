package commentrepository

import (
	"fmt"

	commentmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/comment_services/model"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/google/uuid"
)

func (store *commentStore) CreateNewComment(post_id string, input *commentmodel.CreateComment) error {
	input.PostID = post_id
	input.ID = uuid.New().String()
	input.CreatedAt = utils.GetCurrentDateTime()
	query := `
		insert into comments (id, post_id, user_id, content, created_at)
		values (:id, :post_id, :user_id, :content, :created_at)
	`
	if _, err := store.db.NamedExec(query, input); err != nil {
		return fmt.Errorf("error when insert new comment into db: %w", err)
	}
	return nil
}
