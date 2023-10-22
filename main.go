package main

import (
	"sync"
	"urban-map/api"
	"urban-map/internal"
	"urban-map/internal/utils/db"
	"urban-map/telegram_bot"
)

func main() {
	internal.Init()
	d := db.Connect()
	defer db.Close(d)
	var wg sync.WaitGroup
	wg.Add(2)
	go api.Run(&wg)
	go telegram_bot.Run(&wg)
	wg.Wait()
}
