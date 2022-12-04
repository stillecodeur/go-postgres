package models

type Stock struct {
	StockID int64  `json:"stockID"`
	Name    string `json:"name"`
	Price   int    `json:"price"`
	Company string `json:"company"`
}
