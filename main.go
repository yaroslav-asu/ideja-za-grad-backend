package main

import (
	"urban-map/internal"
	"urban-map/internal/utils/db"
)

func main() {
	internal.Init()
	d := db.Connect()
	defer db.Close(d)

}
