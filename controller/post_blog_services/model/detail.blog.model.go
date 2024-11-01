package blogmodel

type DetailBlogModel struct {
	PostID     string          `db:"post_id" json:"post_id"`
	AuthorName string          `db:"author_name" json:"author_name"`
	Title      string          `db:"title" json:"title"`
	Content    string          `db:"content" json:"content"`
}
