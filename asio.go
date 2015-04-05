package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Asio :)!\n")
}

func DummyMethod(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Dummy!")
}

// Create endpoints for all ASIO functions.
// ASIO must support two versions, the previous and current release.
func CreateRouters(router *httprouter.Router) {
	router.GET("/", Index)
	// *** v1 API REST ***
	// Routes for single domain
	router.GET("/v1/domain", DummyMethod)
	router.GET("/v1/domain/:domain_id", DummyMethod)
	router.POST("/v1/domain", DummyMethod) // Add an domain with settings
	router.PUT("/v1/domain/:domain_id", DummyMethod)
}

func main() {
	router := httprouter.New()
	n := negroni.Classic()
	CreateRouters(router)
	n.UseHandler(router)
	n.Run(":4000")
}
