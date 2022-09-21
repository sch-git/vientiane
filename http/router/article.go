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

	for _, conf := range getRouteConfigs {
		router.GET(conf.Path, conf.Handle)
	}
}

