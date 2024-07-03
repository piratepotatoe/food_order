package main

import (
	"foodorder/internal/controller"
	"foodorder/internal/repo"
	"foodorder/internal/router"
	"foodorder/internal/services"
	"log"
	"net/http"
)

func main() {
	// Initialize OrderRepo once
	orderRepo := repo.NewOrderRepo()

	// Initialize OrderService with the same OrderRepo instance
	orderService := services.NewOrderService(orderRepo)

	// Initialize OrderController with the same OrderService instance
	orderController := controller.NewOrderController(orderService)

	// Set up the router
	r := router.NewRouter(orderController)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
