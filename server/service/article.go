package service

import (
	"context"
	"fmt"
	"vientiane/server/dao"
	"vientiane/server/database"
	"vientiane/server/models"
)

type ArticleService struct {
	db  database.ManagerDB
	dao *dao.ArticleDAO
}

func NewArticleService() models.ArticleService {
	return &ArticleService{
		db:  database.NewDB(),
		dao: dao.NewArticleDAO(),
	}
}

func (a *ArticleService) Get(ctx context.Context, id int64) (*models.Article, error) {
	fun := "ArticleService.Get -->"
	db, err := a.db.GetDB()
	if err != nil {
		err = fmt.Errorf("%s get db err: %v", fun, err)
		return nil, err
	}

	article, err := a.dao.Get(ctx, id, db)
	if err != nil {
		err = fmt.Errorf("%s get article err: %v", fun, err)
		return nil, err
	}
	return article, nil
}

func (a *ArticleService) List(ctx context.Context, article *models.Article) ([]*models.Article, error) {
	panic("implement me")
}

func (a *ArticleService) Add(ctx context.Context, article *models.Article) (err error) {
	panic("implement me")
}

func (a *ArticleService) Del(ctx context.Context, article *models.Article) (err error) {
	panic("implement me")
}
