package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"net/http"
	"time"
	"vientiane/http/result"
)

var Vlog = zap.NewExample()

type Handler interface {
	Handle(ctx *gin.Context)
}

func BindJsonWrapper(h func() Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		defer func() {
			HttpCost.WithLabelValues("test").Observe(time.Now().Sub(start).Seconds())
		}()
		ins := h()
		err := c.Bind(&ins)
		if nil != err {
			Vlog.Error("BindJsonWrapper", zap.Error(err))
			result.RespERR(c, &result.Result{
				Code:    http.StatusBadRequest,
				Message: "bind json err",
			})
			return
		}

		ins.Handle(c)
	}
}

var (
	HttpCost = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "vientiane",
		Name:      "http_cost",
	}, []string{"cost"})
)

func init() {
	prometheus.MustRegister(HttpCost)
}
