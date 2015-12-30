package controller

import (
	// standard library packages
	"encoding/json"
	"net/http"

	// third party packages
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/model"
)

type (
	ApiController struct {
		session *mgo.Session
	}
)

// ApiController apis
func NewApiController(s *mgo.Session) *ApiController {

	return &ApiController{s}
}

func (ac ApiController) Authorize(r *http.Request) bool {

	// Parse Authorization Header
	authorization, err := json.Marshal(r.Header.Get("Authorization"))
	if err != nil {
		// Unauthorized
		return false
	}

	// Strip off Bearer
	token := authorization[7:]

	// Fetch api key
	key := model.Key{}
	if err := ac.session.DB("smartcart").C("keys").Find(bson.M{"value": token}).One(&key); err != nil {
		// Forbidden
		return false
	}

	if key.Value == "" {
		// Forbidden
		return false
	}

	return true
}
