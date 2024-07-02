package router

import (
	c "foodorder/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/menu", c.NewOrderController().GetMenu)
		v1.POST("/order", c.NewOrderController().PlaceOrder)
	}
	return r
}
