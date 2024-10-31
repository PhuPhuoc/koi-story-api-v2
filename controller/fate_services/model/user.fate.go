package fatemodel

type GenerateUserFengshui struct {
	UserID      string `json:"user_id"`
	Gender      string `json:"gender"`
	YearOfBirth int    `json:"year_of_birth"`
}

type UserFengshui struct {
	ID          string `db:"id" json:"-"`
	UserID      string `db:"user_id" json:"-"`
	Gender      string `db:"gender" json:"-"`
	YearOfBirth int    `db:"year_of_birth" json:"year_of_birth"`
	FateID      string `db:"fate_id" json:"-"`
	Direction   string `db:"direction" json:"direction"`
	CungPhi     string `db:"cung_phi" json:"cung_phi"`
}
