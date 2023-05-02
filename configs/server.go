package configs

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/services/utils"
)

func (c *Config) InitServer() (err error) {
	var (
		envPort = utils.GetEnv("APP_PORT", models.APP_PORT)
	)

	port := fmt.Sprintf(":%s", envPort)
	log.Println("SERVER LISTENING ON PORT", port)

	err = http.ListenAndServe(port, c.Router)
	if err != nil {
		return
	}

	return nil
}
