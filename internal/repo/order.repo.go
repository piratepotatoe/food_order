package repo

import (
	"errors"
	"foodorder/models"
	"log"
)

type OrderRepo struct {
	orders []models.Order
	nextID int
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		orders: []models.Order{},
		nextID: 1,
	}
}

var menu = []models.MenuItem{
	{ID: 1, Name: "banh mi", Price: 10.00},
	{ID: 2, Name: "banh cuon", Price: 5.00},
	{ID: 3, Name: "banh bao", Price: 15.00},
}

func (orderRepo *OrderRepo) GetMenu() []models.MenuItem {
	return menu
}

func (orderRepo *OrderRepo) GetMenuByID(id int) (models.MenuItem, error) {
	for _, item := range menu {
		if item.ID == id {
			return item, nil
		}
	}
	return models.MenuItem{}, errors.New("menu item not found")
}

func (orderRepo *OrderRepo) AddOrder(order models.Order) models.Order {
	order.ID = orderRepo.nextID
	orderRepo.nextID++
	orderRepo.orders = append(orderRepo.orders, order)
	log.Println("Order added to slice:", order)
	log.Println("Current orders slice:", orderRepo.orders)
	return order
}

func (orderRepo *OrderRepo) GetPendingOrders() []models.Order {
	var pendingOrders []models.Order
	if len(orderRepo.orders) == 0 {
		log.Println("Empty orders slice:", orderRepo.orders)
		log.Println("Empty pendingOrders slice:", pendingOrders)
		return pendingOrders
	}
	for _, order := range orderRepo.orders {
		log.Println("Checking order:", order)
		if order.Status == "waiting" {
			pendingOrders = append(pendingOrders, order)
			log.Println("Added to pendingOrders:", order)
		}
	}
	log.Println("Final pendingOrders slice:", pendingOrders)
	return pendingOrders
}

func (orderRepo *OrderRepo) CompleteOrder(id int) error {
	for i, order := range orderRepo.orders {
		if order.ID == id {
			orderRepo.orders[i].Status = "completed"
			return nil
		}
	}
	return errors.New("order not found")
}
