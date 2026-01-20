package types

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Sku         string  `json:"sku"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	Image       string  `json:"image"`
}
