package domain

type Product struct {
	ID          int    `json:"id"`
	ItemNumber  string `json:"item_number"`
	Type        string `json:"type"`
	Brand       string `json:"brand"`
	Gender      string `json:"gender"`
	VolumeML    int    `json:"volume_ml"`
	Description string `json:"description"`
	Barcode     string `json:"barcode"`
	CreatedAt   int64  `json:"created_at"`
}
