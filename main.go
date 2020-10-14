package main

import (
	"sync"

	"github.com/p1ck0/broadcast_dir/client"
	util "github.com/p1ck0/broadcast_dir/util"
)

func init() {
	util.CreateDir("broadcast_dir")
}

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	cl := &client.Client{
		Port:         ,
		LocalСlients: []string{},
	}
	go cl.Run()
	go util.Notify(cl)
	wg.Wait()
}
