package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"vientiane/http/handle"
)

func HandleMonitor(r *gin.Engine, moduleName string) {
	router := r.Group(moduleName)
	monitor := prometheus.NewRegistry()
	//promHandler := promhttp.InstrumentMetricHandler(monitor, promhttp.HandlerFor(monitor, promhttp.HandlerOpts{}))
	promHandler := promhttp.InstrumentMetricHandler(monitor, promhttp.Handler())
	postRouteConfigs := []RouteConfig{
		{"", gin.WrapH(promHandler)},
	}

	for _, conf := range postRouteConfigs {
		router.Any(conf.Path, conf.Handle)
	}

	monitor.MustRegister(handle.HttpCost)    // 自定义注册器 promhttp.HandlerFor(monitor, promhttp.HandlerOpts{})
	prometheus.MustRegister(handle.HttpCost) // 使用默认注册器，对应 promhttp.Handler()
}
