package commentmodel

type UpdateComment struct {
	ID      string `db:"id" json:"-"`
	Content string `db:"content" json:"content"`
}
