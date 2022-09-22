package router

import (
	"github.com/gin-gonic/gin"
	"vientiane/http/handle"
	"vientiane/http/handle/article"
)

func HandleArticle(r *gin.Engine, moduleName string) {
	router := r.Group(moduleName)
	getRouteConfigs := []RouteConfig{
		{"/get/:id", handle.BindJsonWrapper(article.FactoryGetArticle)},
	}
	postRouteConfigs := []RouteConfig{
		{"/add", handle.BindJsonWrapper(article.FactoryAddArticle)},
		{"/del", handle.BindJsonWrapper(article.FactoryDelArticle)},
		{"/set", handle.BindJsonWrapper(article.FactorySetArticle)},
	}

	for _, conf := range getRouteConfigs {
		router.GET(conf.Path, conf.Handle)
	}
	for _, conf := range postRouteConfigs {
		router.POST(conf.Path, conf.Handle)
	}
}

