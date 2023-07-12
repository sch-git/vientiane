package router

import (
	"github.com/gin-gonic/gin"
	"vientiane/http/handle"
	"vientiane/http/handle/indices"
)

func HandleIndices(r *gin.Engine, moduleName string) {
	router := r.Group(moduleName)
	postRouteConfigs := []RouteConfig{
		{"/create", handle.BindJsonWrapper(indices.FactoryCreateIndices)},
	}

	//for _, conf := range getRouteConfigs {
	//	router.GET(conf.Path, conf.Handle)
	//}
	for _, conf := range postRouteConfigs {
		router.POST(conf.Path, conf.Handle)
	}
}
