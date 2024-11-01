package consultmodel

type PostConsult struct {
	ConsultModelForGet
	Images   []DetailImage          `json:"images"`
	Comments []CommentInPostConsult `json:"comments"`
}

type ConsultModelForGet struct {
	PostID  string `db:"post_id" json:"post_id"`
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
}

type DetailImage struct {
	PostID   string `db:"post_id" json:"-"`
	ID       string `db:"id" json:"id"`
	ImageUrl string `db:"image_url" json:"image_url"`
}

type CommentInPostConsult struct {
	ID        string `db:"id" json:"id"`
	PostID    string `db:"post_id" json:"-"`
	UserID    string `db:"user_id" json:"user_id"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
	Name      string `db:"name" json:"name"`
	Avatar    string `db:"avatar" json:"avatar"`
}

type CommentModel struct {
}

type UserModel struct {
}
