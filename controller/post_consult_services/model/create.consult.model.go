package model

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
	PostID      string `db:"post_id" json:"-"`
	ProductName string `db:"product_name" json:"product_name"`
	Price       int    `db:"price" json:"price"`
	ProductType string `db:"product_type" json:"product_type"`
	Color       string `db:"color" json:"color"`
	Origin      string `db:"origin" json:"origin"`
	Description string `db:"description" json:"description"`
}
