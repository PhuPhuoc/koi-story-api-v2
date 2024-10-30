package marketrepository

import (
	"database/sql"
	"errors"
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/model"
)

func (store *marketStore) GetPostMarketDetails(post_id string) (obj *marketmodel.ModelDetailMarket, err error) {
	obj = &marketmodel.ModelDetailMarket{}
	obj.ID = post_id
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

	query_post := `
		select user_id, created_at from posts where id=? and deleted_at is null
	`
	if err = tx.Get(&obj.DetailPost, query_post, obj.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("post not found: %w", err)
		}
		return nil, fmt.Errorf("cannot get post's detail: %w", err)
	}

	query_market := `
		select
			product_name, price, product_type, color, origin, description
		from
			markets
		where post_id=?
	`
	if err = tx.Get(&obj.DetailMarket, query_market, obj.ID); err != nil {
		return nil, fmt.Errorf("cannot get market's detail: %w", err)
	}

	query_seller := `
		select phone_number, location, address from seller_info where user_id=?
	`
	if err = tx.Get(&obj.DetailSeller, query_seller, obj.UserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("seller info not found: %w", err)
		}
		return nil, fmt.Errorf("cannot get seller's detail: %w", err)
	}
	query_image := `
		select id, image_url from images where post_id=? and deleted_at is null order by image_order
	`
	if err = tx.Select(&obj.ListImage, query_image, obj.ID); err != nil {
		return nil, fmt.Errorf("cannot get list images of market's post: %w", err)
	}

	return obj, nil
}
