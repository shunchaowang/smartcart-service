package model

import (
    "time"
)

type (
	Order struct {
		Id       int     `json:"id"`
		Amount    float32 `json:"amount"`
		Remark string `json:"remark"`

		CreatedTime time.Time `json:"createdTime"`
		UpdatedTime time.Time `json:"updatedTime"`

		OrderStatus OrderStatus `json:"status"`
		//Customer Customer `json:'customer'`
		Products []Product `json:"products"`

		ShippingAddress Address `json:'address'`
		BillingAddress Address `json:'address'`
	}

	OrderStatus struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}
	
	Address struct {
		Company         string `json:"company"`
		FirstName       string `json:"firstname"`
		LastName        string `json:"lastname"`
		FullName        string `json:"fullname"`
		Street          string `json:"street"`
		StreetLine      string `json:"street_line2"`
		City            string `json:"city"`
		State           string `json:"state"`
		PostCode        string `json:"postcode"`
		Country         string `json:"country"`
	}

)
