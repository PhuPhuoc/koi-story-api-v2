package marketmodel

type PostOfMarketCreation struct {
	ID        string `db:"id" json:"id"`
	UserID    string `db:"user_id" json:"user_id"`
	PostType  string `db:"post_type" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
}

type MarketCreation struct {
	PostID      string `db:"post_id" json:"-"`
	ProductName string `db:"product_name" json:"product_name"`
	Price       int    `db:"price" json:"price"`
	ProductType string `db:"product_type" json:"product_type"`
	Color       string `db:"color" json:"color"`
	Origin      string `db:"origin" json:"origin"`
	Description string `db:"description" json:"description"`
}

type FormCreatePostMarket struct {
	PostOfMarketCreation
	MarketCreation
	ListImageUrls []string
}
