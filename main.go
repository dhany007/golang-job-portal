package main

import (
	"encoding/json"
	"net/http"

	"github.com/dhany007/golang-job-portal/configs"
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
	app := configs.New()

	configs.Catch(app.InitService())
	configs.Catch(app.InitEnv())
	configs.Catch(app.InitPostgres())
	configs.Catch(app.InitServer())

	app.Start()
}
