package store

type Product struct {
	ID     int64   `bson:"_id"`
	Title  string  `json:"title"`
	Image  string  `json:"image"`
	Price  uint64  `json:"price"`
	Rating float32 `json:"rating"`
}

type Products []Product
