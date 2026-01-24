package types

type Product struct {
	Sku         string `json:"sku" db:"sku"`
	Title       string `json:"title" db:"title"`
	Brand       string `json:"brand" db:"brand"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
	Rating      int    `json:"rating" db:"rating"`
	ImageUrl    string `json:"image_url" db:"image_url"`
}

type Products []Product
