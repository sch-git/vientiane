package router

import (
	"github.com/gin-gonic/gin"
	"vientiane/http/handle"
	"vientiane/http/handle/account"
)

func HandleAccount(r *gin.Engine, moduleName string) {
	router := r.Group(moduleName)
	getRouteConfigs := []RouteConfig{
		{"/get/:id", handle.BindJsonWrapper(account.FactoryGetAccount)},
		{"/list", handle.BindJsonWrapper(account.FactoryListAccount)},
	}
	postRouteConfigs := []RouteConfig{
		//{"/list", handle.BindJsonWrapper(account.FactoryListAccount)},
	}

	for _, conf := range getRouteConfigs {
		router.GET(conf.Path, conf.Handle)
	}
	for _, conf := range postRouteConfigs {
		router.POST(conf.Path, conf.Handle)
	}
}
