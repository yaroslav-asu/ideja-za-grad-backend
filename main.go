package main

import (
	"log"
	"urban-map/api"
	"urban-map/internal"
	"urban-map/internal/utils/db"
)

func main() {
	internal.Init()
	d := db.Connect()
	defer db.Close(d)
	router := api.InitRouter()
	err := router.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
