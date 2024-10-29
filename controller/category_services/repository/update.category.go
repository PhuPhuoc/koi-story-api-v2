package categoryrepository

import (
	"fmt"

	categorymodel "github.com/PhuPhuoc/koi-story-api-v2/controller/category_services/model"
	"github.com/jmoiron/sqlx"
)

func (store *categoryStore) UpdateCategory(cate_id string, data *categorymodel.CategoryForCreateAndUpdate) (err error) {
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

	data.ID = cate_id
	query_update_cate := `
	    update categories set name=:name, description=:description where id=:id
	`
	if _, err = tx.NamedExec(query_update_cate, data); err != nil {
		return fmt.Errorf("cannot update category data: %w", err)
	}

	list_fate_id := []string{}
	query_get_fate := `
        select fate_id from category_fate where category_id=?
    `
	if err = tx.Select(&list_fate_id, query_get_fate, cate_id); err != nil {
		return fmt.Errorf("cannot get fate belong to category to update: %w", err)
	}
	toDelete := checkDifferent(data.FateID, list_fate_id)
	toAdd := checkDifferent(list_fate_id, data.FateID)

	if len(toDelete) > 0 {
		var (
			query_delete string
			args         []any
		)
		query_delete, args, err = sqlx.In("delete from category_fate where category_id=? and fate_id in (?)", cate_id, toDelete)
		if err != nil {
			return fmt.Errorf("cannot delete old fate belong to category: %w", err)
		}
		_, err = tx.Exec(query_delete, args...)
		if err != nil {
			return fmt.Errorf("cannot exec query to delete old fate: %w", err)
		}
	}

	if len(toAdd) > 0 {
		for _, fateID := range toAdd {
			_, err = tx.Exec("insert into category_fate (category_id, fate_id) VALUES (?, ?)", cate_id, fateID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func checkDifferent(slice1, slice2 []string) []string {
	new_map := make(map[string]bool)
	for _, item := range slice1 {
		new_map[item] = true
	}

	diff := []string{}
	for _, item := range slice2 {
		if _, found := new_map[item]; !found {
			diff = append(diff, item)
		}
	}
	return diff
}
