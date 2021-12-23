package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"vientiane/server/models"
)

type AccountDAO struct {
}

func NewAccountDAO() *AccountDAO {
	return &AccountDAO{}
}

func (d *AccountDAO) Get(ctx context.Context, id int64, db *gorm.DB) (account *models.Account, err error) {
	fun := "AccountDAO.Get-->"

	result := db.First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		account = &models.Account{}
		glog.Infof("%s get account by id: %d err: not found", fun, id)
		return
	}

	if result.Error != nil {
		err = fmt.Errorf("%s get account by id: %d err: %v", fun, id, result.Error)
		return
	}

	return
}
