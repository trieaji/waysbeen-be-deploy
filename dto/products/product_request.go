package productsdto

type CreateProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"`
	Stock int    `json:"stock"`
	Desc  string `json:"desc"`
}
