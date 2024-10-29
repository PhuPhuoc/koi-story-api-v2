package categoryrepository

import (
	"fmt"

	categorymodel "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/model"
	"github.com/google/uuid"
)

func (store *categoryStore) CreateCategory(data *categorymodel.CategoryForCreateAndUpdate) (err error) {
	tx, err := store.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p) // Rethrow panic after rollback.
		} else if err != nil {
			_ = tx.Rollback() // Ignore rollback error and return original error.
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction: %w", commitErr)
		}
	}()

	new_cate_id := uuid.New().String()
	cate := categorymodel.CategoriesModel{
		ID:          new_cate_id,
		Name:        data.Name,
		Description: data.Description,
	}

	query_cate := `
		insert into categories (id, name, description)
		values (:id, :name, :description)
	`
	if _, err = tx.NamedExec(query_cate, cate); err != nil {
		return fmt.Errorf("cannot insert new category: %w", err)
	}

	query_cate_fate := `
		insert into category_fate (category_id, fate_id)
		values (?,?)
	`
	for _, fid := range data.FateID {
		if _, err = tx.Exec(query_cate_fate, new_cate_id, fid); err != nil {
			return fmt.Errorf("cannot insert fate into category_fate: %w", err)
		}
	}

	return nil
}
