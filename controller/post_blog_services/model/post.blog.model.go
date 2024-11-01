package blogmodel

type PostBlogModel struct {
	PostID     string          `db:"post_id" json:"post_id"`
	AuthorName string          `db:"author_name" json:"author_name"`
	Title      string          `db:"title" json:"title"`
}
