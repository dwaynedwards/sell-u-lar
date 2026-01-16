package types

import "time"

type Product struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Category    string   `json:"category"`
	Price       float64  `json:"price"`
	Rating      float64  `json:"rating"`
	Tags        []string `json:"tags"`
	Brand       string   `json:"brand"`
	Sku         string   `json:"sku"`
	Reviews     []Review `json:"reviews"`
	Images      []string `json:"images"`
	Thumbnail   string   `json:"thumbnail"`
}

type Review struct {
	Rating        int       `json:"rating"`
	Comment       string    `json:"comment"`
	Date          time.Time `json:"date"`
	ReviewerName  string    `json:"reviewerName"`
	ReviewerEmail string    `json:"reviewerEmail"`
}
