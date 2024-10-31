package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/model"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/google/uuid"
)

func (store *consultStore) CreateNewConsultPost(input *consultmodel.CreateConsultModel) (err error) {
	tx, err := store.db.Beginx()
	if err != nil {
		return fmt.Errorf("cannot begin transaction: %w", err)
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

	postID := uuid.New().String()
	input.ID = postID
	input.PostType = "consult"
	input.CreatedAt = utils.GetCurrentDateTime()
	input.PostID = postID

	queryPost := `
		insert into posts (id, user_id, post_type, created_at)
		values (:id, :user_id, :post_type, :created_at)
	`
	if _, err = tx.NamedExec(queryPost, input.PostOfConsultModel); err != nil {
		return fmt.Errorf("cannot create new post: %w", err)
	}

	queryMarket := `
		insert into consults (post_id, title , content)
		values (:post_id, :title , :content)
	`
	if _, err = tx.NamedExec(queryMarket, input.ConsultModel); err != nil {
		return fmt.Errorf("cannot create detail market: %w", err)
	}

	queryImage := `
		insert into images (id, post_id, image_url, image_order)
		values (?, ?, ?, ?)
	`
	for index, url := range input.ListImageUrls {
		imgID := uuid.New().String()
		if _, err = tx.Exec(queryImage, imgID, postID, url, index); err != nil {
			return fmt.Errorf("cannot insert image into post: %w", err)
		}
	}
	return nil
}
