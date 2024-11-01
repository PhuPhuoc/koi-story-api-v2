package blogrepository

import (
	"fmt"

	blogmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/model"
)

func (store *blogStore) GetBlogByCategory(category_id string) ([]blogmodel.PostBlogModel, error) {
	blogs := []blogmodel.PostBlogModel{}

	query := `
		select post_id, author_name, title
		from blogs
		where category_id=?
	`
	if err := store.db.Select(&blogs, query, category_id); err != nil {
		return nil, fmt.Errorf("unable to select blogs: %w", err)
	}
	return blogs, nil
}
