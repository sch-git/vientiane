package handle

import (
	"github.com/gin-gonic/gin"
	. "vientiane/http/router"
	"vientiane/http/router/health"
)

func HandleHealth(r *gin.Engine, moduleName string) {
	router := r.Group(moduleName)
	getRouteConfigs := []RouteConfig{
		{"/health", health.HealthCheckIn()},
	}

	for _, conf := range getRouteConfigs {
		router.GET(conf.Path, conf.Handle)
	}
}
