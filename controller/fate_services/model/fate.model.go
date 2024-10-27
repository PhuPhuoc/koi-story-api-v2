package fatemodel

type Fate struct {
	ID       string `db:"id" json:"id"`
	FateName string `db:"name" json:"fate_name"`
}
