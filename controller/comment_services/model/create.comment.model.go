package commentmodel

type CreateComment struct {
	ID        string `db:"id" json:"-"`
	PostID    string `db:"post_id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"-"`
}
