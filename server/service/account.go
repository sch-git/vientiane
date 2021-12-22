package service

import (
	"vientiane/server/dao"
	"vientiane/server/db"
	"vientiane/server/models"
)

type AccountService struct {
	db  db.ManagerDB
	dao *dao.AccountDAO
}

func NewAccountService() *AccountService {
	return &AccountService{
		db:  db.NewDB(),
		dao: dao.NewAccountDAO(),
	}
}

func (s *AccountService) Get(id int64) (account *models.Account, err error) {
	if id < 1 {
		return
	}

	return
}
