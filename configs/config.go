package configs

import (
	"github.com/dhany007/golang-job-portal/services/repository/database"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Router      *httprouter.Router
	DB          *database.DB
	RedisClient *redis.Client
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
