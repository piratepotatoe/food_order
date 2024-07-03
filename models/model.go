package models

type MenuItem struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type OrderRequest struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}
type Order struct {
	ID       int    `json:"id"`
	ItemID   int    `json:"item_id"`
	Quantity int    `json:"quantity"`
	Status   string `json:"status"` // eg. waiting, completed
}
