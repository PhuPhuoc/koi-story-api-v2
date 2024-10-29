package usermodel

type SellerInfo struct {
	UserID      string `db:"user_id" json:"-"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	Location    string `db:"location" json:"location"`
	Address     string `db:"address" json:"address"`
}
