package controller

import (
	"foodorder/internal/services"
	"foodorder/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (orderController *OrderController) PlaceOrder(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("Placing order:", orderReq)
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

func (orderController *OrderController) GetPendingOrders(c *gin.Context) {
	orders := orderController.orderService.GetPendingOrders()
	log.Println("Pending orders retrieved:", orders)
	c.JSON(http.StatusOK, gin.H{"pendingOrders": orders})
}

func (orderController *OrderController) CompleteOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := orderController.orderService.CompleteOrder(order.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order completed successfully"})
}
