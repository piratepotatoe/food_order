package services

import (
	"fmt"
	"foodorder/internal/repo"
	"foodorder/models"
	"log"
)

type OrderService struct {
	orderRepo *repo.OrderRepo
}

func NewOrderService(orderRepo *repo.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (orderService *OrderService) GetMenu() []models.MenuItem {
	return orderService.orderRepo.GetMenu()
}

func (orderService *OrderService) PlaceOrder(itemID, quantity int) (string, error) {
	item, err := orderService.orderRepo.GetMenuByID(itemID)
	if err != nil {
		return "", err
	}
	totalAmount := item.Price * float64(quantity)
	order := models.Order{
		ItemID:   itemID,
		Quantity: quantity,
		Status:   "waiting",
	}
	log.Println("Adding order:", order)
	orderService.orderRepo.AddOrder(order)
	return fmt.Sprintf("₹%.2f", totalAmount), nil
}

func (orderService *OrderService) GetPendingOrders() []models.Order {
	pendingOrders := orderService.orderRepo.GetPendingOrders()
	log.Println("Pending orders in service:", pendingOrders)
	return pendingOrders
}

func (orderService *OrderService) CompleteOrder(id int) error {
	return orderService.orderRepo.CompleteOrder(id)
}
