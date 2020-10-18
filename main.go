package main

import (
	"sync"

	"github.com/p1ck0/broadcast_dir/client"
	"github.com/p1ck0/broadcast_dir/util"
)

func init() {
	util.CreateDir("broadcast_dir")
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	cl := &client.Client{
		Port:         "1488",
		LocalСlients: []string{"127.0.0.1:1489"},
	}
	go cl.Run()
	go util.Notify(cl)
	wg.Wait()
}
