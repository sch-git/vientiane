package service

import (
	"context"
	"fmt"
	"vientiane/server/dao"
	"vientiane/server/database"
	"vientiane/server/models"
)

type AccountService struct {
	db  database.ManagerDB
	dao *dao.AccountDAO
}

func NewAccountService() models.AccountService {
	return &AccountService{
		db:  database.NewDB(),
		dao: dao.NewAccountDAO(),
	}
}

func (s *AccountService) Get(ctx context.Context, id int64) (account *models.Account, err error) {
	fun := "AccountService.Get-->"
	if id < 1 {
		return
	}

	db, err := s.db.GetDB()
	if nil != err {
		err = fmt.Errorf("%s get db err: %v", fun, err)
		return
	}

	account, err = s.dao.Get(ctx, id, db)
	if nil != err {
		err = fmt.Errorf("%s %v", fun, err)
		return
	}

	return
}

func (s *AccountService) List(ctx context.Context, account *models.Account) (accounts []*models.Account, err error) {
	fun := "AccountService.List-->"

	if nil == account {
		err = fmt.Errorf("%s account is nil", fun)
		return
	}

	db, err := s.db.GetDB()
	if nil != err {
		err = fmt.Errorf("%s get db err: %v", fun, err)
		return
	}

	accounts, err = s.dao.List(ctx, account, db)
	if nil != err {
		err = fmt.Errorf("%s %v", fun, err)
		return
	}

	return
}
