package test

import (
	"github.com/golang/glog"
	"testing"
	"vientiane/server/dao"
	. "vientiane/server/db"
)

func TestGet(t *testing.T) {
	d := NewDB()
	db, err := d.GetDB()
	if nil != err {
		glog.Errorf("get db err: %v", err)
		return
	}

	accountDAO := dao.NewAccountDAO()
	account, err := accountDAO.Get(1, db)
	if nil != err {
		glog.Errorf("get account err: %v", err)
		return
	}

	t.Log("res", account)
}
