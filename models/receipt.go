package models

// Receipt
type Receipt struct {
	retailer     string
	purchaseDate string
	purchaseTime string
	items        []Item
	points       int
}

func NewReceipt(r string, pd string, pt string, i []Item, p int) Receipt {
	return Receipt{retailer: r, purchaseDate: pd, purchaseTime: pt, items: i, points: p}
}

// Item
type Item struct {
	shortDescription string
	price            string
}

func NewItem(shortDescription string, price string) Item {
	return Item{shortDescription: shortDescription, price: price}
}
