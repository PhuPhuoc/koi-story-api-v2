package fatemodel

type GenerateUserFengshui struct {
	UserID      string `json:"user_id"`
	YearOfBirth int    `json:"year_of_birth"`
}

type UserFengshui struct {
	ID            string `db:"id" json:"-"`
	UserID        string `db:"user_id" json:"-"`
	YearOfBirth   int    `db:"year_of_birth" json:"year_of_birth"`
	FateID        string `db:"fate_id" json:"-"`
	HeavenlyStem  string `db:"heavenly_stem" json:"heavenly_stem"`
	EarthlyBranch string `db:"earthly_branch" json:"earthly_branch"`
}
