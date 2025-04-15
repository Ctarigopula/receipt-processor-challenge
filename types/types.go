package types

// Receipt struct
type Receipt struct {
	ID           string `json:"id"`
	Retailer     string `json:"retailer"`
	PurchaseTime string `json:"purchaseTime"`
	PurchaseDate string `json:"purchaseDate"`
	Total        string `json:"total"`
	Items        []Item `json:"items"`
}
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
