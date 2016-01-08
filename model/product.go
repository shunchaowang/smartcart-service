package model

type (
	Product struct {
		Id       int     `json:"id"`
		Model    string  `json:"model"`
		Quantity int     `json:"quantity"`
		Image    string  `json:"image"`
		Price    float32 `json:"price"`

		Weight float32 `json:"weight"`

		Name        string `json:"name"`
		Description string `json:"description"`

		Category Category `json:"category"`
	}

	Category struct {
		Id    int    `json:"id"`
		Image string `json:"image"`

		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
