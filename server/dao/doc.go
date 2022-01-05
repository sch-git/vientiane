package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"vientiane/server/models"
)

type DocDAO struct {
}

func NewDocDAO() *DocDAO {
	return &DocDAO{}
}

func (d *DocDAO) Add(ctx context.Context, doc *models.Doc, db *gorm.DB) (err error) {
	fun := "DocDAO.Add -->"
	if nil == db || nil == doc {
		err = fmt.Errorf("%s db: %v or doc:%v is nil", fun, db, doc)
		return
	}

	err = db.Create(doc).Error
	if nil != err {
		err = fmt.Errorf("%s add doc err: %v", fun, err)
		return
	}

	return
}
