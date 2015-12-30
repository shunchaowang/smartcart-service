package model

import "gopkg.in/mgo.v2/bson"

type (
	Key struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Name  string        `json:"name" bson:"name"`
		Value string        `json:"value" bson:"value"`
	}
)
