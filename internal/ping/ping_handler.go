package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/constant"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/utils"
)

type PingHandler struct {
	r  *gin.Engine
	ps Ping
}

func ProvidePingHandler(ps Ping, r *gin.Engine) PingHandler {
	return PingHandler{
		ps: ps,
		r:  r,
	}
}

func (p *PingHandler) InitHandler() {
	pingRouter := p.r.Group("/api/v1")
	{
		pingRouter.Handle(http.MethodGet, constant.PING_API_PATH, p.Ping)
	}
}

func (p *PingHandler) Ping(c *gin.Context) {
	res := p.ps.Ping()
	response := utils.NewSuccessResponseWriter(c.Writer, http.StatusOK, "SUCCESS", res)
	c.JSON(http.StatusOK, response)
}
