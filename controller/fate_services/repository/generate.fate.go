package faterepository

import (
	"fmt"

	fatemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/model"
	"github.com/google/uuid"
)

func (store *fateStore) GenerateFates() error {
	var count int
	query := `SELECT COUNT(*) FROM fates`
	if err := store.db.Get(&count, query); err != nil {
		return fmt.Errorf("cannot count fates: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("fate has enough, no need to create more")
	}

	fates := []fatemodel.Fate{
		{ID: uuid.New().String(), FateName: "Metal"},
		{ID: uuid.New().String(), FateName: "Wood"},
		{ID: uuid.New().String(), FateName: "Water"},
		{ID: uuid.New().String(), FateName: "Fire"},
		{ID: uuid.New().String(), FateName: "Earth"},
	}

	insertQuery := `INSERT INTO fates (id, name) VALUES (:id, :name)`
	if _, err := store.db.NamedExec(insertQuery, fates); err != nil {
		return fmt.Errorf("cannot generate fate: %w", err)
	}
	return nil
}
