package imagerepository

import (
	"fmt"

	imagemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/image_services/model"
	"github.com/google/uuid"
)

func (store *imageStore) AddNewImage(post_id string, input *imagemodel.AddImageModel) error {
	var current_order int
	query_get_current_order := `
		select
			max(image_order)
		from
			images
		where
			post_id=?
	`
	if err := store.db.Get(&current_order, query_get_current_order, post_id); err != nil {
		return fmt.Errorf("cannot get current orther of list image in the post: %w", err)
	}

	queryImage := `
		insert into images (id, post_id, image_url, image_order)
		values (?, ?, ?, ?)
	`
	current_order++
	imgID := uuid.New().String()
	if _, err := store.db.Exec(queryImage, imgID, post_id, input.ImageUrl, current_order); err != nil {
		return fmt.Errorf("cannot insert new image into post: %w", err)
	}
	return nil
}
