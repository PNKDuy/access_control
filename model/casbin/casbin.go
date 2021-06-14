package casbin

import "access_control/model"

type Casbin struct {
	PType string `json:"p_type"`
	V0    string `json:"v0"`
	V1    string `json:"v1"`
	V2    string `json:"v2"`
}

const (
	selectCasbinByRole = "select p_type, v0, v1, v2 from casbins where v0 = ? LIMIT 1000"
)

func GetByRole(role string) ([]Casbin, error) {
	var casbins []Casbin
	if err := model.Db.Raw(selectCasbinByRole, role).Scan(&casbins).Error; err != nil {
		return nil, err
	}
	return casbins, nil
}

func (casbin Casbin)Create() (Casbin,error) {
	if err := model.Db.Create(&casbin).Error; err != nil {
		return casbin, err
	}
	return casbin, nil
}

func Delete(role, path, method string) error {
	if err := model.Db.Where("v0 = ? && v1 = ? && v2 = ?", role, path, method).Delete(&Casbin{}).Error; err != nil {
		return err
	}
	return nil
}