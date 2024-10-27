package faterepository

import (
	"fmt"

	fatemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/model"
	"github.com/google/uuid"
)

func (store *fateStore) CalculateUserFate(data *fatemodel.GenerateUserFengshui) error {
	can, chi, menh := getZodiacElement(data.YearOfBirth)
	var fate_id string
	if err := store.db.Get(&fate_id, `select id from fates where name=?`, menh); err != nil {
		return fmt.Errorf("cannot get fate in db: %w", err)
	}

	user_fengshui := fatemodel.UserFengshui{
		ID:            uuid.New().String(),
		UserID:        data.UserID,
		YearOfBirth:   data.YearOfBirth,
		FateID:        fate_id,
		HeavenlyStem:  can,
		EarthlyBranch: chi,
	}

	query := `
	insert into user_fengshui (id, user_id, year_of_birth, fate_id, heavenly_stem, earthly_branch)
	values (:id, :user_id, :year_of_birth, :fate_id, :heavenly_stem, :earthly_branch)
	`
	_, err := store.db.NamedExec(query, user_fengshui)
	if err != nil {
		return fmt.Errorf("failed to create new user: %w", err)
	}
	return nil
}

func getZodiacElement(year int) (string, string, string) {
	// Danh sách thiên can và địa chi
	canList := []string{"Canh", "Tân", "Nhâm", "Quý", "Giáp", "Ất", "Bính", "Đinh", "Mậu", "Kỷ"}
	chiList := []string{"Tý", "Sửu", "Dần", "Mão", "Thìn", "Tỵ", "Ngọ", "Mùi", "Thân", "Dậu", "Tuất", "Hợi"}

	// Danh sách mệnh ngũ hành theo can-chi
	elementList := []string{"Metal", "Metal", "Water", "Water", "Wood", "Wood", "Fire", "Fire", "Earth", "Earth"}

	// Tính chỉ số thiên can và địa chi từ năm sinh
	can := year % 10
	chi := (year - 4) % 12

	// Xác định mệnh ngũ hành từ chỉ số thiên can
	element := elementList[can]

	// Trả về thiên can, địa chi, và mệnh
	return canList[can], chiList[chi], element
}
