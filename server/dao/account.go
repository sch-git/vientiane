package dao

import (
	"errors"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"vientiane/server/models"
)

type AccountDAO struct {
}

func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

func (d *AccountDAO) Get(id int64, db *gorm.DB) (account *models.Account, err error) {
	fun := "AccountDAO.Get-->"

	result := db.First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		account = &models.Account{}
		glog.Infof("%s get account by id: %d err: not found", fun, id)
		return
	}

	if result.Error != nil {
		err = result.Error
	}

	return
}
