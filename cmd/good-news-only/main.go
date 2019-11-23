package main

import (
	"fmt"
	"runtime"

	cron "github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 2s", func() { fmt.Println("Hello") })
	c.Start()
	runtime.Goexit()
	c.Stop()
}
