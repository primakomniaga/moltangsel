package http

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jemmycalak/mall-tangsel/api"
	"github.com/jemmycalak/mall-tangsel/api/http/product"
	"github.com/jemmycalak/mall-tangsel/api/http/user"
	u "github.com/jemmycalak/mall-tangsel/api/utils"
)

type Server struct {
	server         *http.Server
	UserService    api.UserService
	ProductService api.ProductService
}

func (s *Server) Serve(lis net.Listener) error {
	s.server = &http.Server{}

	// init all handler
	user.Init(s.UserService)
	product.Init(s.ProductService)

	// import all route into server handler
	s.server.Handler = Handler()

	return s.server.Serve(lis)
}

func Handler() *gin.Engine {

	c := gin.Default()
	c.Use(u.ApikeyValidator())
	c.MaxMultipartMemory = 1 << 2

	v1 := c.Group("/api/v1")
	{
		//user
		v1.POST("/account/register", user.Register)
		v1.POST("/account/login", user.Login)

		//product
		v1.GET("/product/:id", product.Product)
		v1.GET("/products", product.Products)
		v1.GET("/products/limit", product.ProductLimit)

		token := v1.Group("")
		token.Use(u.TokenValidator())
		{
			//user
			token.GET("/account", user.GetUser)

			//product
			token.PUT("/product", product.EditProduct)
			token.DELETE("/product/:id", product.DeleteProduct)
			token.PUT("/product/image/:id", product.UpdateImage)
			token.POST("/product", product.NewProduct)
		}
	}

	c.NoRoute(func(c *gin.Context) {
		u.ResponseError(c, http.StatusBadGateway, "I dont know what are you looking for !")
		c.Abort()
		return
	})
	return c
}
