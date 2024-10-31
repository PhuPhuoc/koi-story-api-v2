package marketrepository

import (
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/model"
)

func (store *marketStore) UpdateMarket(post_id string, data *marketmodel.MarketInfoUpdate) error {
	data.PostID = post_id
	query := `
		update
			markets
		set
			product_name=:product_name, price=:price, product_type=:product_type, color=:color, origin=:origin, description=:description
		where post_id=:post_id
	`

	if _, err := store.db.NamedExec(query, data); err != nil {
		return fmt.Errorf("cannot update market: %w", err)
	}
	return nil
}
