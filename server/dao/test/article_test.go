package test

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"log"
	"testing"
	"vientiane/server/dao"
	. "vientiane/server/database"
	"vientiane/server/models"
)

var (
	articleDB  *gorm.DB
	articleDAO *dao.ArticleDAO
	articleCtx context.Context
	err        error
)

func init() {
	d := NewDB()
	articleDB, err = d.GetDB()
	if nil != err {
		glog.Errorf("get db err: %v", err)
		return
	}
	articleDAO = dao.NewArticleDAO()
}

func TestAddArticle(t *testing.T) {
	article := &models.Article{
		Title:   "title",
		Content: "content",
		Author:  "author",
	}
	go func() {
		err = articleDAO.Add(articleCtx, article, articleDB)
		if err != nil {
			t.Logf("error -->")
		}
	}()
}

func TestGetArticle(t *testing.T) {
	go func() {
		_, err = articleDAO.Get(articleCtx, 1, articleDB)
		if err != nil {
			log.Println(err)
		}
	}()
}
