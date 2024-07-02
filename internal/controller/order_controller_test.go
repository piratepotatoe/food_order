package controller

import (
	"bytes"
	"encoding/json"
	"foodorder/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetMenu(t *testing.T) {
	router := gin.Default()
	orderController := NewOrderController()
	router.GET("/menu", orderController.GetMenu)

	req, _ := http.NewRequest("GET", "/menu", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string][]models.MenuItem
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.NotEmpty(t, response["menu"])
}

func TestPlaceOrder(t *testing.T) {
	router := gin.Default()
	orderController := NewOrderController()
	router.POST("/order", orderController.PlaceOrder)

	orderReq := models.OrderRequest{
		ItemID:   1,
		Quantity: 2,
	}
	jsonValue, _ := json.Marshal(orderReq)
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Order placed successfully", response["message"])
	assert.NotEmpty(t, response["totalAmount"])
}
