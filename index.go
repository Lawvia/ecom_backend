package main

import (
	"fmt"
	"ecom_backend/app/mappings"
	"ecom_backend/app/modules"
)

func main() {
	new(modules.Configuration).Init()
	config := new(modules.Configuration).GetConfig()
	port := fmt.Sprintf(":%v",config.GetString("APP_PORT"))

	db := new(modules.DbConnect)
	db.Init()

	mappings.CreateUrlMappings()
	mappings.Router.Run(port)

	defer db.GetDB().Close()
}
