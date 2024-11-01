package categoryrepository

import (
	"fmt"

	categorymodel "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/model"
)

func (store *categoryStore) GetCategories() ([]categorymodel.CategoriesDisplay, error) {
	response := []categorymodel.CategoriesDisplay{}
	res_cate := []categorymodel.CategoriesModel{}

	query_select_cate := `
		select id, name, description from categories
	`
	if err := store.db.Select(&res_cate, query_select_cate); err != nil {
		return nil, fmt.Errorf("cannot get list categories: %w", err)
	}

	for _, cate := range res_cate {
		res_fate := []categorymodel.CateroryFate{}

		query_select_cate := `
			select f.id, f.element from fates f join category_fate cf on cf.fate_id=f.id where cf.category_id=?
		`
		if err := store.db.Select(&res_fate, query_select_cate, cate.ID); err != nil {
			return nil, fmt.Errorf("cannot get list categories: %w", err)
		}

		res_data := categorymodel.CategoriesDisplay{
			CategoriesModel: cate,
			Fate:            res_fate,
		}
		response = append(response, res_data)
	}

	return response, nil
}
