package consultmodel

type CreateConsultModel struct {
	PostOfConsultModel
	ConsultModel
	ListImageUrls []string
}

type PostOfConsultModel struct {
	ID        string `db:"id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	PostType  string `db:"post_type" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
}

type ConsultModel struct {
	PostID  string `db:"post_id" json:"-"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}
