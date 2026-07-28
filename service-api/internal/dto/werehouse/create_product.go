package werehouse

type CreateProductRequest struct {
	ItemNumber string      `json:"itemNumber"`
	Type       string      `json:"type"`
	Brand      string      `json:"brand"`
	Meta       ProductMeta `json:"meta"`
}

type ProductMeta struct {
	Gender      string `json:"gender"`
	VolumeML    int    `json:"volumeMl"`
	Description string `json:"description"`
	Barcode     string `json:"barcode"`
}
