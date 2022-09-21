package dao

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"vientiane/server/models"
)

type ArticleDAO struct {
}

func NewArticleDAO() *ArticleDAO {
	return &ArticleDAO{}
}

func (d *ArticleDAO) Add(ctx context.Context, article *models.Article, db *gorm.DB) (err error) {
	fun := "ArticleDAO.Add-->"
	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	if nil == article || article.IsEmpty() {
		err = fmt.Errorf("%s article is nil or article is empty", fun)
		return
	}

	err = db.Create(article).Error
	if nil != err {
		err = fmt.Errorf("%s %v", fun, err)
		return
	}

	return
}

func (d *ArticleDAO) Get(ctx context.Context, id int64, db *gorm.DB) (article *models.Article, err error) {
	fun := "ArticleDAO.Get-->"

	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	result := db.Where("id=?", id).First(&article)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		article = &models.Article{}
		return
	}

	if result.Error != nil {
		err = fmt.Errorf("%s get article by id: %d err: %v", fun, id, result.Error)
		return
	}

	return
}

func (d *ArticleDAO) List(ctx context.Context, article *models.Article, db *gorm.DB) (articles []*models.Account, err error) {
	fun := "ArticleDAO.List-->"
	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	if nil == article {
		err = fmt.Errorf("%s article is nil", fun)
		return
	}

	db = db.Offset(int(article.Offset)).Limit(int(article.Limit))

	err = db.Find(&articles).Error
	if nil != err {
		err = fmt.Errorf("%s list article by req: %v err: %v", fun, article, err)
		return
	}

	return
}

func (d *ArticleDAO) Set(ctx context.Context, article *models.Article, db *gorm.DB) (err error) {
	fun := "ArticleDAO.Set-->"
	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	if nil == article || article.Id < 1 {
		err = fmt.Errorf("%s article is nil or id is empty", fun)
		return
	}

	data := map[string]interface{}{
		"title":   article.Title,
		"content": article.Content,
	}
	where := map[string]interface{}{
		"id": article.Id,
	}

	err = db.Model(&article).Where(where).Updates(data).Error
	if nil != err {
		err = fmt.Errorf("%s update article err: %v", fun, err)
		return
	}

	return
}

func (d *ArticleDAO) Count(ctx context.Context, article *models.Article, db *gorm.DB) (count int64, err error) {
	fun := "ArticleDAO.Count-->"
	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	db = db.Model(article)
	err = db.Count(&count).Error
	if nil != err {
		err = fmt.Errorf("%s %v", fun, err)
		return
	}

	return
}

func (d *ArticleDAO) Del(ctx context.Context, id int64, db *gorm.DB) (err error) {
	fun := "ArticleDAO.Del -->"
	if db == nil {
		err = fmt.Errorf("%s db is nil", fun)
		return
	}

	db = db.Delete(&models.Article{Id: id})
	err = db.Error
	if nil != err {
		err = fmt.Errorf("%s %v",fun,err)
		return
	}

	return
}
