package router

import (
	c "foodorder/internal/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(orderController *c.OrderController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/menu", orderController.GetMenu)
		v1.POST("/orders", orderController.PlaceOrder)
		v1.GET("/orders/pending", orderController.GetPendingOrders)
		v1.POST("/orders/complete", orderController.CompleteOrder)
	}
	return r
}
