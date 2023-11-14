package main

import (
	"github.com/yaroslav-asu/urban-map/api"
	"github.com/yaroslav-asu/urban-map/internal"
	"github.com/yaroslav-asu/urban-map/internal/utils/db"
	"github.com/yaroslav-asu/urban-map/telegram_bot"
	"sync"
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
