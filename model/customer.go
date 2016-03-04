package model

type (
	Customer struct {
		Id       int     `json:"id"`
		Gender    string  `json:"gender"`
		FirstName string     `json:"firstName"`
        LastName string `json:"lastName"`
        Dob string `json:"dob"`
        Email string `json:"email"`
        Password string `json:"password"`

        Address Address `json:"address"`
	}

    Address struct {
        Recipient string `json:"recipient"`
        Tel string `json:"tel"`
        Street string `json:"street"`
        City string `json:"city"`
        State string `json:"state"`
        Zip string `json:"zip"`
        Country string `json:"country"`

    }
)
