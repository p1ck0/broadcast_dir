package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func (client *Client) BroadCast(filename string) {
	var pack = packageTCP{
		Filename: filename,
	}
	for _, clientIP := range client.LocalСlients {
		wg.Add(1)
		go func(clientIP string) {
			defer wg.Done()
			file, err := os.Open("broadcast_dir/" + filename)
			if err != nil {
				fmt.Println(err)
			}
			conn, err := net.Dial("tcp", clientIP)
			if err != nil {
				fmt.Println(err)
				return
			}
			sum := client.Files[filename]
			pack.SHA256 = sum
			jsonpack, _ := json.Marshal(pack)
			if err != nil {
				fmt.Println(err)
			}
			conn.Write(jsonpack)
			var buf = make([]byte, 32*1024)
			n, err := io.CopyBuffer(conn, file, buf)
			fmt.Println(filename, n)
			file.Close()
			conn.Close()
		}(clientIP)
	}
	wg.Wait()
}
