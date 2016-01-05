package model

type (
	Product struct {
		Id    int `json:"id"`
		Model string        `json:"model"`
        Type ProductType `json:"type"`
	}

	ProductType struct {
		Id   int `json:"id"`
		Name string        `json:"name"`
	}
)
