package commentmodel

type Comment struct {
	CommentModel
	UserModel
}

type CommentModel struct {
	ID        string `db:"id" json:"id"`
	UserID    string `db:"user_id" json:"user_id"`
	Content   string `db:"content" json:"content"`
	CreatedAt string `db:"created_at" json:"created_at"`
}

type UserModel struct {
	Name   string `db:"name" json:"name"`
	Avatar string `db:"avatar" json:"avatar"`
}
