package marketrepository

import (
	"fmt"

	marketmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_market_services/model"
)

func (store *marketStore) GetMyPostMarketList(user_id string) ([]marketmodel.PostMarket, error) {
	list := []marketmodel.PostMarket{}
	query := `
		select
			m.post_id, m.product_name, m.price, m.product_type, i.image_url
		from
			markets m
		join
			images i on i.post_id=m.post_id
		join
			posts p on p.id=m.post_id
		WHERE
			p.deleted_at is null
			and
			i.image_order in (select min(image_order) from images where i.post_id=m.post_id)
			and
			i.deleted_at is null
			and
			p.user_id=?
		group by
			m.post_id, m.product_name, m.price, m.product_type, i.image_url
		order by
			p.created_at desc
	`
	if err := store.db.Select(&list, query, user_id); err != nil {
		return nil, fmt.Errorf("cannot select post market from db: %w", err)
	}
	return list, nil
}
