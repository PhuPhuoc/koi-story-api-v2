package consultmodel

type ModelUpdateConsult struct {
	PostID  string `db:"post_id" json:"-"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}
