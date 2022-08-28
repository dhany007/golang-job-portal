package configs

import (
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Router *httprouter.Router
	DB     *database.DB
}

func New() Config {
	var cfg Config
	return cfg
}

func (c *Config) Start() (err error) {
	ch := make(chan bool)

	go func() {
		err = c.Start()
		if err != nil {
			return
		}
	}()

	<-ch
	return nil
}

func (c *Config) InitEnv() (err error) {
	if err = godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func Catch(err error) {
	if err != nil {
		panic(err)
	}
}
