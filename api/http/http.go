package http

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/mall-tangsel/api"
	"github.com/jemmycalak/mall-tangsel/api/http/user"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
)

type Server struct {
	server      *http.Server
	UserService api.UserService
}

func (s *Server) Serve(lis net.Listener) error {
	s.server = &http.Server{}

	// init all handler
	user.Init(s.UserService)

	// import all route into server handler
	s.server.Handler = Handler()

	return s.server.Serve(lis)
}

func Handler() *gin.Engine {
	c := gin.Default()
	v1 := c.Group("/api/v1/account")
	{
		v1.POST("/", user.Register)
		v1.GET("/", user.GetUser)

	}
	c.NoRoute(func(c *gin.Context) {
		u.ResponseError(c, http.StatusBadGateway, "I dont know what are you looking for !")
		c.Abort()
		return
	})
	return c
}
