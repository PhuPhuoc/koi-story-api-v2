package marketrepository

import (
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/model"
	"github.com/PhuPhuoc/koi-story-api-v2/utils"
	"github.com/google/uuid"
)

func (store *marketStore) CreateNewPost(input *marketmodel.FormCreatePostMarket) (err error) {
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
	input.PostType = "market"
	input.CreatedAt = utils.GetCurrentDateTime()
	input.PostID = postID

	queryPost := `
		insert into posts (id, user_id, post_type, created_at)
		values (:id, :user_id, :post_type, :created_at)
	`
	if _, err = tx.NamedExec(queryPost, input.PostOfMarketCreation); err != nil {
		return fmt.Errorf("cannot create new post: %w", err)
	}

	queryMarket := `
		insert into markets (post_id, product_name, price, product_type, color, origin, description)
		values (:post_id, :product_name, :price, :product_type, :color, :origin, :description)
	`
	if _, err = tx.NamedExec(queryMarket, input.MarketCreation); err != nil {
		return fmt.Errorf("cannot create detail market: %w", err)
	}

	queryImage := `
		insert into images (id, post_id, image_url, image_order)
		values (?, ?, ?, ?)
	`
	for index, url := range input.ListImageUrls {
		imgID := uuid.New().String()
		if _, err = tx.Exec(queryImage, imgID, postID, url, index); err != nil {
			return fmt.Errorf("cannot insert image in post: %w", err)
		}
	}
	return nil
}
