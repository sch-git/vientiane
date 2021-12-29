package test

import (
	"context"
	"github.com/golang/glog"
	"testing"
	"vientiane/server/dao"
	. "vientiane/server/database"
	"vientiane/server/models"
)

func TestGet(t *testing.T) {
	d := NewDB()
	db, err := d.GetDB()
	if nil != err {
		glog.Errorf("get db err: %v", err)
		return
	}

	accountDAO := dao.NewAccountDAO()
	account, err := accountDAO.Get(context.TODO(), 1, db)
	if nil != err {
		glog.Errorf("get account err: %v", err)
		return
	}

	t.Log("res", account)
}

func TestList(t *testing.T) {
	d := NewDB()
	db, err := d.GetDB()
	if nil != err {
		glog.Errorf("get db err: %v", err)
		return
	}

	accountDAO := dao.NewAccountDAO()
	accounts, err := accountDAO.List(context.TODO(), &models.Account{Name: "", Limit: 10}, db)
	if nil != err {
		glog.Errorf("list account err: %v", err)
		return
	}

	t.Log("res", accounts)
}
