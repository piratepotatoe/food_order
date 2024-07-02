package controller

import (
	"foodorder/internal/services"
	"foodorder/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController() *OrderController {
	return &OrderController{
		orderService: services.NewOrderService(),
	}
}
func (orderController *OrderController) PlaceOrder(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBind(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	totalAmount, err := orderController.orderService.PlaceOrder(orderReq.ItemID, orderReq.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "totalAmount": totalAmount})
}

func (orderController *OrderController) GetMenu(c *gin.Context) {
	menu := orderController.orderService.GetMenu()
	c.JSON(http.StatusOK, gin.H{"menu": menu})
}
