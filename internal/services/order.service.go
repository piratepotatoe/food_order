package services

import (
	"fmt"
	"foodorder/internal/repo"
	"foodorder/models"
)

type OrderService struct {
	orderRepo *repo.OrderRepo
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderRepo: repo.NewOrderRepo(),
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
	return fmt.Sprintf("₹%.2f", totalAmount), nil
}
