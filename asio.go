package main

import (
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func CreateRouters(router *httprouter.Router) {
	router.GET("/", Index)
}

func main() {
	router := httprouter.New()
	n := negroni.Classic()
	CreateRouters(&router)
	n.UseHandler(router)
	n.Run(":4000")
}
