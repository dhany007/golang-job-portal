package main

import (
	"github.com/dhany007/golang-job-portal/configs"
)

func main() {
	app := configs.New()

	configs.Catch(app.InitEnv())
	configs.Catch(app.InitRedis())
	configs.Catch(app.InitPostgres())
	configs.Catch(app.InitService())
	configs.Catch(app.InitServer())

	app.Start()
}
