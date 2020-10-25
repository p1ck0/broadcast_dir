package client

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func (client *Client) BroadCast(filename string) {
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
			}
			sum := client.Files[filename]
			_, err = io.WriteString(conn, fmt.Sprintf("%s\n", string(sum[:])))
			if err != nil {
				fmt.Println(err)
			}
			_, err = io.WriteString(conn, fmt.Sprintf("%s\n", filename))
			if err != nil {
				fmt.Println(err)
			}
			n, err := io.Copy(conn, file)
			fmt.Println(n)
			file.Close()
			conn.Close()
		}(clientIP)
	}
	wg.Wait()
}
