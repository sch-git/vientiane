package dao

import (
	"context"
	"errors"
	"fmt"
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

	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	result := db.First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		account = &models.Account{}
		//glog.Infof("%s get account by id: %d err: not found", fun, id)
		return
	}

	if result.Error != nil {
		err = fmt.Errorf("%s get account by id: %d err: %v", fun, id, result.Error)
		return
	}

	return
}

func (d *AccountDAO) List(ctx context.Context, account *models.Account, db *gorm.DB) (accounts []*models.Account, err error) {
	fun := "AccountDAO.List-->"
	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	if nil == account {
		err = fmt.Errorf("%s account is nil", fun)
		return
	}

	if account.Name != "" {
		db = db.Where("name like ?", "%"+account.Name+"%")
	}
	if account.Email != "" {
		db = db.Where("email = ?", account.Email)
	}
	db = db.Offset(int(account.Offset)).Limit(int(account.Limit))

	err = db.Find(&accounts).Error
	if nil != err {
		err = fmt.Errorf("%s list account by req: %v err: %v", fun, account, err)
		return
	}

	return
}
