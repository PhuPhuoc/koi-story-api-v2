package blogrepository

import (
	"fmt"

	blogmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/model"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/google/uuid"
)

func (store *blogStore) CreateNewBlogPost(input *blogmodel.CreationBlogModel) (err error) {
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
	input.PostType = "blog"
	input.CreatedAt = utils.GetCurrentDateTime()
	input.PostID = postID

	queryPost := `
		insert into posts (id, user_id, post_type, created_at)
		values (:id, :user_id, :post_type, :created_at)
	`
	if _, err = tx.NamedExec(queryPost, input.PostOfBlogModel); err != nil {
		return fmt.Errorf("cannot create new post: %w", err)
	}

	queryBlog := `
		insert into blogs (post_id, category_id, author_name, title , content)
		values (:post_id, :category_id, :author_name, :title , :content)
	`
	if _, err = tx.NamedExec(queryBlog, input.BlogModel); err != nil {
		return fmt.Errorf("cannot create detail market: %w", err)
	}
	return nil
}
