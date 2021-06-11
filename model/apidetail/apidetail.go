package apidetail

import (
	"access_control/model"
	"time"

	"github.com/google/uuid"
)


type ApiDetail struct {
	Id        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Path 		string `json:"path"`
	Method		string	`json:"method"`
	Description	string	`json:"description"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	IsDeleted bool      `json:"-"`
}

const (
	selectApiDetailFalseDeleted = "SELECT id, path, method, description FROM api_details where is_deleted is false "
	deactiveApiDetail = "UPDATE api_details SET is_deleted = true WHERE id = ?"
)

func Get() ([]ApiDetail, error){
	var apiDetails []ApiDetail
	if err := model.Db.Raw(selectApiDetailFalseDeleted + " LIMIT 100").Scan(&apiDetails).Error; err != nil {
		return nil, err
	}
	return apiDetails, nil
}

func GetById(id string) (ApiDetail, error) {
	var apiDetail = ApiDetail{}
	if err := model.Db.Raw(selectApiDetailFalseDeleted+ "AND id = ?", id).Scan(&apiDetail).Error; err != nil {
		return apiDetail, err
	}
	return apiDetail, nil
}

func (apiDetail ApiDetail)Create()(ApiDetail, error) {
	apiDetail.Id = uuid.New()
	apiDetail.UpdatedAt = time.Now()
	apiDetail.CreatedAt = time.Now()
	apiDetail.IsDeleted = false
	if err := model.Db.Create(&apiDetail).Error; err != nil {
		return apiDetail, err
	}
	return apiDetail, nil
}

func (apiDetail ApiDetail)Update() (ApiDetail, error){
	if err := model.Db.Model(&apiDetail).Update(apiDetail).Error; err != nil {
		return apiDetail, err
	}
	return apiDetail, nil
}

func Delete(id string) error {
	if err := model.Db.Exec(deactiveApiDetail, id).Error; err != nil {
		return err
	}
	return nil
}