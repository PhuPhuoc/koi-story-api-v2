package blogmodel

type CreationBlogModel struct {
	PostOfBlogModel
	BlogModel
}

type PostOfBlogModel struct {
	ID        string `db:"id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	PostType  string `db:"post_type" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
}

type BlogModel struct {
	PostID     string `db:"post_id" json:"-"`
	CategoryID string `db:"category_id" json:"category_id"`
	AuthorName string `db:"author_name" json:"author_name"`
	Title      string `db:"title" json:"title"`
	Content    string `db:"content" json:"content"`
}
