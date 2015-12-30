package model

import "gopkg.in/mgo.v2/bson"

type (
	Product struct {
		Id    bson.ObjectId `json:"id"`
		Model string        `json:"model"`
	}

	ProductType struct {
		Id   bson.ObjectId `json:"id"`
		Name string        `json:"name"`
	}
)
