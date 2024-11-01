package categorymodel

type CategoriesModel struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

type CateroryFate struct {
	ID   string `db:"id" json:"id"`
	Element string `db:"element" json:"element"`
}

type CategoriesDisplay struct {
	CategoriesModel
	Fate []CateroryFate
}

type CategoryForCreateAndUpdate struct {
	ID          string   `json:"-"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	FateID      []string `json:"fate_id"`
}
