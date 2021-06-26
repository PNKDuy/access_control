package role

import (
	"access_control/model"
	"time"

	"github.com/google/uuid"
)


type Role struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDeleted bool      `json:"-"`
}

const (
	selectRoleFalseDeleted = "select id, role_name from roles where is_deleted = false "
	updateRoleById = "UPDATE roles SET is_deleted = true Where id = ?"
)

func Get() ([]Role, error){
	var roles []Role
	if err := model.Db.Raw(selectRoleFalseDeleted + "LIMIT 100").Scan(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func GetById(id string) (Role, error) {
	var role = Role{}
	if err := model.Db.Raw(selectRoleFalseDeleted+ "AND id = ?", id).Scan(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (role Role)Create()(Role, error) {
	if role.Id == uuid.Nil {
		role.Id = uuid.New()
	}
	role.UpdatedAt = time.Now()
	role.CreatedAt = time.Now()
	role.IsDeleted = false
	if err := model.Db.Create(&role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func (role Role)Update() (Role, error){
	if err := model.Db.Model(&role).Update(role).Error; err != nil {
		return role, err
	}
	return role, nil
}

func Delete(id string) error {
	if err := model.Db.Exec(updateRoleById, id).Error; err != nil {
		return err
	}
	return nil
}