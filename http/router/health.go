package router

import (
	"github.com/gin-gonic/gin"
	"vientiane/http/handle"
	"vientiane/http/handle/health"
)

func HandleHealth(r *gin.Engine, moduleName string) {
	router := r.Group(moduleName)
	getRouteConfigs := []RouteConfig{
		{"/ping", handle.BindJsonWrapper(health.HealthCheck)},
	}

	for _, conf := range getRouteConfigs {
		router.GET(conf.Path, conf.Handle)
	}
}
