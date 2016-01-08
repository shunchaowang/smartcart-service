package main

import (
	// standard library packages
	"database/sql"
	"fmt"
	"net/http"

	// third party packages
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	//zencartcontroller "github.com/shunchaowang/zencart-service/controller"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/controller"
)

func main() {

	// Http request should contain a hreader with authorization: Bearer <uid>
	// Server will check uid against backend mongo db
	// instantiate a new router
	// two routers listen on different ports
	zencartrouter := httprouter.New()
	magentorouter := httprouter.New()

	// /zencart will be routed to zencart service
	// /magento will be routed to magento service

	// get a ProductController instance

	// Verify api key

	// get a product
	zencartrouter.GET("/zencart", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("Listening on zencart."))
	})

	magentorouter.GET("/magento", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("Listening on magento."))
	})

	// create a product

	// delete a product

	go func() {
		http.ListenAndServe("localhost:8081", magentorouter)
	}()

	// fire up the server
	http.ListenAndServe("localhost:8080", zencartrouter)
}

// TODO: TBD
func handler(w http.ResponseWriter, r *http.Request) {
	// check api key
	ac := controller.NewApiController(getMongoSession())
	if !ac.Authorize(r) {
		w.WriteHeader(403) // Forbidden
	}

	// get service type

}

func authorize(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println(*r.URL)
		ac := controller.NewApiController(getMongoSession())
		if !ac.Authorize(r) {
			w.WriteHeader(403)
			return
		}
		fn(w, r, p)
	}
}

// Create a new mongo session and panics if connection error occurs
func getMongoSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection fails, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}

func getZencartDB() *sql.DB {
	db, err := sql.Open("mysql", "zencart:zencart@tcp(localhost:3306)/zencart?charset=utf8")

	if err != nil {
		panic(err)
	}

	return db
}
