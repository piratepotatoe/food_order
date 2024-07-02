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
