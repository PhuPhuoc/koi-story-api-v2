package fatemodel

type UserFate struct {
	YearOfBirth   int    `db:"year_of_birth" json:"year_of_birth"`
	FateName      string `db:"name" json:"fate_name"`
	HeavenlyStem  string `db:"heavenly_stem" json:"heavenly_stem"`
	EarthlyBranch string `db:"earthly_branch" json:"earthly_branch"`
}
