package consultmodel

type PostConsult struct {
	PostID   string `db:"post_id" json:"post_id"`
	Title    string `db:"title" json:"title"`
	Content  string `db:"content" json:"content"`
	ImageUrl string `db:"image_url" json:"image_url"`
}
