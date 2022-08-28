package configs

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (c *Config) InitServer() (err error) {
	var (
		envPort = os.Getenv("APP_PORT")
	)

	port := fmt.Sprintf(":%s", envPort)
	log.Println("SERVER LISTENING ON PORT", port)

	err = http.ListenAndServe(port, c.Router)
	if err != nil {
		return
	}

	return nil
}
