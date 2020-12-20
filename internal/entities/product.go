package entities

// Product is an entity that represents a product in the system
type Product struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	Title       string `json:"Title"`
}
