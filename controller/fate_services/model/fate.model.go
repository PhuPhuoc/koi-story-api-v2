package fatemodel

type Fate struct {
	ID      string `db:"id" json:"id"`
	Element string `db:"element" json:"element"`
}
