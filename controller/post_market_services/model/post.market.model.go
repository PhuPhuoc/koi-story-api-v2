package marketmodel

type PostMarket struct {
	PostID      string `db:"post_id" json:"post_id"`
	ProductName string `db:"product_name" json:"product_name"`
	Price       int    `db:"price" json:"price"`
	ProductType string `db:"product_type" json:"product_type"`
	ImageUrl    string `db:"image_url" json:"image_url"`
}
