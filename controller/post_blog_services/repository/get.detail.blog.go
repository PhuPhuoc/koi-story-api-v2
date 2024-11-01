package blogrepository

import (
	"fmt"

	blogmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/model"
)

func (store *blogStore) GetBlogDetail(post_id string) (*blogmodel.DetailBlogModel, error) {
	blogs := &blogmodel.DetailBlogModel{}

	query := `
		select post_id, author_name, title, content
		from blogs
		where post_id=?
	`
	if err := store.db.Get(blogs, query, post_id); err != nil {
		return nil, fmt.Errorf("unable to get blogs: %w", err)
	}
	return blogs, nil
}
