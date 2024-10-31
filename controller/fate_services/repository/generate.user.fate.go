package faterepository

import (
	"fmt"

	fatemodel "github.com/PhuPhuoc/koi-story-api-v2/controller/fate_services/model"
	"github.com/google/uuid"
)

func (store *fateStore) CalculateUserFate(data *fatemodel.GenerateUserFengshui) error {
	var existingID string
	if err := store.db.Get(&existingID, `select id from user_fengshui where user_id = ?`, data.UserID); err == nil {
		if _, err := store.db.Exec(`delete from user_fengshui where user_id = ?`, data.UserID); err != nil {
			return fmt.Errorf("failed to delete existing user_fengshui record: %w", err)
		}
	}

	phongthuy := TinhPhongThuy(data.Gender, data.YearOfBirth)
	var fate_id string
	if err := store.db.Get(&fate_id, `select id from fates where element=?`, phongthuy.menh); err != nil {
		return fmt.Errorf("cannot get fate in db: %w", err)
	}

	user_fengshui := fatemodel.UserFengshui{
		ID:          uuid.New().String(),
		UserID:      data.UserID,
		Gender:      data.Gender,
		YearOfBirth: data.YearOfBirth,
		FateID:      fate_id,
		Direction:   phongthuy.huong,
		CungPhi:     phongthuy.cungphi,
	}

	query := `
	insert into user_fengshui (id, user_id, gender, year_of_birth, fate_id, direction, cung_phi)
	values (:id, :user_id, :gender, :year_of_birth, :fate_id, :direction, :cung_phi)
	`
	_, err := store.db.NamedExec(query, user_fengshui)
	if err != nil {
		return fmt.Errorf("failed to create new user: %w", err)
	}
	return nil
}

type PhongThuy struct {
	cungphi string
	huong   string
	menh    string
}

func TongChuSo(namAm int) int {
	tong := 0
	for namAm > 0 {
		tong += namAm % 10
		namAm /= 10
	}
	return tong
}

func TraCungPhi(gioiTinh string, soDu int) string {
	bangCungPhi := map[int]map[string]string{
		1: {"male": "Khảm", "female": "Cấn"},
		2: {"male": "Ly", "female": "Càn"},
		3: {"male": "Cấn", "female": "Đoài"},
		4: {"male": "Đoài", "female": "Cấn"},
		5: {"male": "Càn", "female": "Ly"},
		6: {"male": "Khôn", "female": "Khảm"},
		7: {"male": "Tốn", "female": "Khôn"},
		8: {"male": "Chấn", "female": "Chấn"},
		9: {"male": "Khôn", "female": "Tốn"},
	}
	return bangCungPhi[soDu][gioiTinh]
}

func XacDinhMenhHuong(cungPhi string) (string, string) {
	menhHuong := map[string]struct {
		Menh  string
		Huong string
	}{
		"Khảm": {"Thủy", "Bắc"},
		"Ly":   {"Hỏa", "Nam"},
		"Cấn":  {"Thổ", "Đông Bắc"},
		"Đoài": {"Kim", "Tây"},
		"Càn":  {"Kim", "Tây Bắc"},
		"Khôn": {"Thổ", "Tây Nam"},
		"Tốn":  {"Mộc", "Đông Nam"},
		"Chấn": {"Mộc", "Đông"},
	}
	return menhHuong[cungPhi].Menh, menhHuong[cungPhi].Huong
}

func TinhPhongThuy(gioiTinh string, namAm int) PhongThuy {
	tongChuSo := TongChuSo(namAm)

	soDu := tongChuSo % 9
	if soDu == 0 {
		soDu = 9
	}

	cungPhi := TraCungPhi(gioiTinh, soDu)

	menh, huong := XacDinhMenhHuong(cungPhi)

	return PhongThuy{
		cungphi: cungPhi,
		huong:   huong,
		menh:    menh,
	}
}
