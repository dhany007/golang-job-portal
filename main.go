package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handlerTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Add("content-type", "application/json")
	response := map[string]string{
		"message": "success",
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}

func main() {
	router := httprouter.New()

	router.GET("/ping", handlerTest)

	err := http.ListenAndServe(":11010", router)
	if err != nil {
		panic(err)
	}
}
