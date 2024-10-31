package fatemodel

type UserFate struct {
	YearOfBirth int    `db:"year_of_birth" json:"year_of_birth"`
	Element     string `db:"element" json:"element"`
	Direction   string `db:"direction" json:"direction"`
	CungPhi     string `db:"cung_phi" json:"cung_phi"`
}
