package main

import (
    util "github.com/p1ck0/broadcast_dir/util"
    "github.com/p1ck0/broadcast_dir/client"
)

func init() {
    util.CreateDir("broadcast_dir")
}


func main() {
    cl := &client.Client{
        Port : "1488",
    }
    cl.Run()
}
