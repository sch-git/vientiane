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

func TestAdd(t *testing.T) {
	d := NewDB()
	db, err := d.GetDB()
	if nil != err {
		glog.Errorf("get db err: %v", err)
		return
	}

	accountDAO := dao.NewAccountDAO()
	err = accountDAO.Add(context.TODO(), &models.Account{Name: "", Password: "123", Email: "@.com"}, db)
	if nil != err {
		glog.Errorf("add account err: %v", err)
		return
	}

	t.Log("end")
}

type Ac struct {
	Name string
}

func (a *Ac) Empty() bool {
	if a == nil {
		return false
	}
	return *a == Ac{}
}

//type Bc struct {
//	Names []string
//}
//
//func (a Bc) Empty() bool {
//	return a == Bc{}
//}
func TestEmpty(t *testing.T) {
	var a = &Ac{}
	//var b = &Bc{}
	t.Log(a.Empty())
	//t.Log(b.Empty())
}
