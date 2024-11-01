package blogmodel

type DetailBlogModel struct {
	PostID     string          `db:"post_id" json:"post_id"`
	AuthorName string          `db:"author_name" json:"author_name"`
	Title      string          `db:"title" json:"title"`
	Content    string          `db:"content" json:"content"`
	Comments   []CommentInBlog `json:"comments"`
}

type CommentInBlog struct {
	ID        string `db:"id" json:"id"`
	PostID    string `db:"post_id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
	Name      string `db:"name" json:"name"`
	Avatar    string `db:"avatar" json:"avatar"`
}
