package client

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"net"
	"os"
)

func (client *Client) BroadCast(filename string) {
	file, err := os.Open("broadcast_dir/" + filename)
	if err != nil {
		fmt.Println(err)
	}
	s := bufio.NewScanner(file)
	sum := sha256.Sum256(s.Bytes())
	client.files[filename] = sum
	go func() {
		defer file.Close()
		for _, clientIP := range client.LocalСlients {
			conn, _ := net.Dial("tcp", clientIP)
			defer conn.Close()
			_, err := io.WriteString(conn, fmt.Sprintf("%s\n", string(sum[:])))
			if err != nil {
				fmt.Println(err)
			}
			_, err = io.WriteString(conn, fmt.Sprintf("%s\n", filename))
			if err != nil {
				fmt.Println(err)
			}
			_, err = io.Copy(conn, file)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
}
