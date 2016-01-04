package model

type (
	Product struct {
		Id    int `json:"id"`
		Model string        `json:"model"`
	}

	ProductType struct {
		Id   int `json:"id"`
		Name string        `json:"name"`
	}
)
