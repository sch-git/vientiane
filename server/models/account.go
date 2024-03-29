package models

import (
	"context"
	"time"
	pub "vientiane/pub/idl/grpc"
	"vientiane/server/consts"
)

const accountTableName = "vientiane_account"

type AccountService interface {
	Get(ctx context.Context, id int64) (*Account, error)
	List(cxt context.Context, account *Account) ([]*Account, error)
}

type Account struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Limit     int64     `json:"limit" gorm:"-"`
	Offset    int64     `json:"offset" gorm:"-"`
}

func (m *Account) TableName() string {
	return accountTableName
}

func (m *Account) ToGrpc() *pub.Account {
	account := &pub.Account{}
	if nil == m || m.IsEmpty() {
		return account
	}

	return &pub.Account{
		Id:        m.Id,
		Name:      m.Name,
		Password:  m.Password,
		Email:     m.Email,
		UpdatedAt: m.UpdatedAt.Format(consts.TimeFormatLayout),
		CreatedAt: m.CreatedAt.Format(consts.TimeFormatLayout),
	}
}

func (m Account) IsEmpty() bool {
	return m == Account{}
}
