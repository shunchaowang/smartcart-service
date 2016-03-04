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
        Customer Customer `json:'customer'`
        Products []Product `json:"products"`

        ShippingAddress Address `json:'address'`
	}

    OrderStatus struct {
        Id int `json:"id"`
        Name string `json:"name"`
    }

)
