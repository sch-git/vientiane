package test

import (
	"context"
	"github.com/golang/glog"
	"testing"
	"vientiane/server/dao"
	. "vientiane/server/database"
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
