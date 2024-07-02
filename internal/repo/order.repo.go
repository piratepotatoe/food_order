package repo

import (
	"errors"
	"foodorder/models"
)

type OrderRepo struct {
}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
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
