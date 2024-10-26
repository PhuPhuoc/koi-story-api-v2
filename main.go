package main

import (
	"log"

	"github.com/PhuPhuoc/koi-story-api-v2/api"
	"github.com/PhuPhuoc/koi-story-api-v2/db/mysql"
)

//	@title		Koi Story API
//	@version	1.0

func main() {
	db, appport := mysql.ConnectDB()
	server_port := ":" + appport
	sv := api.InitServer(server_port, db)
	if err := sv.RunApp(); err != nil {
		log.Fatal("server (main): ", err)
	}
}
