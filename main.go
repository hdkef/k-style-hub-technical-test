package main

import (
	"fmt"
	"kstylehub/app/config"
	"kstylehub/app/driver"
	"kstylehub/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	//load config
	config := config.NewAPPConfig()
	//init db
	db := driver.InitDB(config)

	r := gin.Default()
	//init routes
	routes.New(r, db, config)

	//start listening
	r.Run(fmt.Sprintf(":%s", config.APP_PORT))
}
