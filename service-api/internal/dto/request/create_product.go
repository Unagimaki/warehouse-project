package request

type CreateProductRequest struct {
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	Gender      string `json:"gender"`
	VolumeML    int    `json:"volumeMl"`
	Description string `json:"description"`
	Barcode     string `json:"barcode"`
}
