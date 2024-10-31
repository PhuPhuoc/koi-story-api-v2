package marketmodel


type MarketInfoUpdate struct {
	PostID      string `db:"post_id" json:"-"`
	ProductName string `db:"product_name" json:"product_name"`
	Price       int    `db:"price" json:"price"`
	ProductType string `db:"product_type" json:"product_type"`
	Color       string `db:"color" json:"color"`
	Origin      string `db:"origin" json:"origin"`
	Description string `db:"description" json:"description"`
}
