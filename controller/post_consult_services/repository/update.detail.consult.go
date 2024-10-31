package consultrepository

import (
	"fmt"

	consultmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_consult_services/model"
)

func (store *consultStore) UpdateConsult(post_id string, data *consultmodel.ModelUpdateConsult) error {
	data.PostID = post_id
	query := `
		update
			consults
		set
			title=:title, content=:content
		where post_id=:post_id
	`

	if _, err := store.db.NamedExec(query, data); err != nil {
		return fmt.Errorf("cannot update consult: %w", err)
	}
	return nil
}
