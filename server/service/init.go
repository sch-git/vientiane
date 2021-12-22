package service

import "vientiane/server/models"

var (
	AccountServiceImpl models.AccountService
)

func init() {
	AccountServiceImpl = NewAccountService()
}
