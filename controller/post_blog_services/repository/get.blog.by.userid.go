package blogrepository

import (
	"fmt"

	blogmodel "github.com/PhuPhuoc/koi-story-api-v2/controller/post_blog_services/model"
)

func (store *blogStore) GetBlogByUser(user_id string) ([]blogmodel.PostBlogModel, error) {
	blogs := []blogmodel.PostBlogModel{}

	query := `
		SELECT b.post_id, b.author_name, b.title
		FROM blogs b
		JOIN categories c ON b.category_id = c.id
		JOIN category_fate cf ON c.id = cf.category_id
		JOIN user_fengshui uf ON uf.fate_id = cf.fate_id
		WHERE uf.user_id=?
	`
	if err := store.db.Select(&blogs, query, user_id); err != nil {
		return nil, fmt.Errorf("unable to select blogs: %w", err)
	}
	return blogs, nil
}
