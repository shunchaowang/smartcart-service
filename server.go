package main

import (
	// standard library packages
	"fmt"
	"net/http"

	// third party packages
	"github.com/julienschmidt/httprouter"
	zencartcontroller "github.com/shunchaowang/zencart-service/controller"
	"gopkg.in/mgo.v2"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/controller"
)

func main() {

	// Http request should contain a hreader with authorization: Bearer <uid>
	// Server will check uid against backend mongo db
	// instantiate a new router
	r := httprouter.New()

	// /zencart will be routed to zencart service
	// /magento will be routed to magento service

	// get a ProductController instance

	// Verify api key

	// get a product
	pc := zencartcontroller.NewProductController(getSession())
	r.GET("/test", authorize(pc.Get))

	// create a product

	// delete a product

	// fire up the server
	http.ListenAndServe("localhost:8080", r)
}

// TODO: TBD
func handler(w http.ResponseWriter, r *http.Request) {
	// check api key
	ac := controller.NewApiController(getSession())
	if !ac.Authorize(r) {
		w.WriteHeader(403) // Forbidden
	}

	// get service type

}

func authorize(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println(*r.URL)
		ac := controller.NewApiController(getSession())
		if !ac.Authorize(r) {
			w.WriteHeader(403)
			return
		}
		fn(w, r, p)
	}
}

// Create a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection fails, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
