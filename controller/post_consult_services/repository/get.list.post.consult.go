package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/model"
)

func (store *consultStore) GetListPostConsult() ([]consultmodel.PostConsult, error) {
	list := []consultmodel.PostConsult{}
	query := `
		select
			c.post_id, c.title, c.content, i.image_url
		from
			consults c
		join
			images i on i.post_id=c.post_id
		join
			posts p on p.id=c.post_id
		WHERE
			p.deleted_at is null
			and
			i.image_order in (select min(image_order) from images where i.post_id=c.post_id)
			and
			i.deleted_at is null
		group by
			c.post_id, c.title, c.content, i.image_url
		order by
			p.created_at desc
	`
	if err := store.db.Select(&list, query); err != nil {
		return nil, fmt.Errorf("cannot select post market from db: %w", err)
	}
	return list, nil
}
