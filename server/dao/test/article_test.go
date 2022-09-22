package test

import (
	"context"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"log"
	"sync"
	"testing"
	"time"
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
	start := time.Now()
	wg:=sync.WaitGroup{}
	for i:=0;i<10000;i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err = articleDAO.Get(articleCtx, 1, articleDB)
			if err != nil {
				log.Println(err)
			}
		}()
	}
	wg.Wait()
	log.Println(time.Now().Sub(start))
}
