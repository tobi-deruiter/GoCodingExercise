package models

// DATA MODELS
// Receipt
type Receipt struct {
	id           string
	points       int
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

func (r Receipt) Id() string {
	return r.id
}

func (r Receipt) Points() int {
	return r.points
}

func (r Receipt) SetId(id string) {
	r.id = id
}

func (r Receipt) SetPoints(points int) {
	r.points = points
}

// Item
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Variables
var receipts []Receipt // list of all reciepts

// Receipt list handlers
func GetReceipts() []Receipt {
	return receipts
}

func AddReceipt(r Receipt) {
	receipts = append(receipts, r)
}
