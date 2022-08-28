package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h handler) PingTest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.testUsecase.PingTest()

	// return response
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
