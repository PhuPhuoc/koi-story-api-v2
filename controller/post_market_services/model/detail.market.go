package marketmodel

type ModelDetailMarket struct {
	ID string `db:"id" json:"id"`
	DetailPost
	DetailMarket
	DetailSeller
	ListImage []DetailImage
}

type DetailPost struct {
	UserID    string `db:"user_id" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
}

type DetailMarket struct {
	ProductName string `db:"product_name" json:"product_name"`
	Price       int    `db:"price" json:"price"`
	ProductType string `db:"product_type" json:"product_type"`
	Color       string `db:"color" json:"color"`
	Origin      string `db:"origin" json:"origin"`
	Description string `db:"description" json:"description"`
}

type DetailSeller struct {
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Location    string `db:"location" json:"location"`
	Address     string `db:"address" json:"address"`
}

type DetailImage struct {
	ID       string `db:"id" json:"id"`
	ImageUrl string `db:"image_url" json:"image_url"`
}
