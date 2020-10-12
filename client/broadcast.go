package client

import (
	"encoding/json"
	"io/ioutil"
	"net"
)

func (client *Client) BroadCast(filename *string) {
	file, _ := ioutil.ReadFile("broadcast_dir/" + *filename)
	var pack = &packageTCP{
		From: head{
			Filename: *filename,
		},
		Body: file,
	}
	data, _ := json.Marshal(pack)
	go func() {
		for _, clientIP := range client.local_clients {
			conn, _ := net.Dial("tcp", clientIP)
			defer conn.Close()
			conn.Write(data)
		}
	}()
}
