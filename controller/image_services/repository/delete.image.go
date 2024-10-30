package imagerepository

import (
	"fmt"

	"github.com/PhuPhuoc/koi-story-api-v2/utils"
)

func (store *imageStore) DeleteImage(image_id string) error {
	date := utils.GetCurrentDateTime()
	query := `
		update images set deleted_at=? where id=?
	`
	if _, err := store.db.Exec(query, date, image_id); err != nil {
		return fmt.Errorf("error when delete image: %w", err)
	}
	return nil
}
