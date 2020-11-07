package client

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

func (client *Client) reciveConn(conn *net.Conn) {
	var (
		pack *packageTCP
	)
	reader := bufio.NewReader(*conn)
	bytepack, _ := reader.ReadBytes('}')
	json.Unmarshal(bytepack, &pack)
	if checkSum := client.Files[pack.Filename]; !bytes.Equal(checkSum[:], pack.SHA256[:]) {
		fmt.Println("checksum: ", checkSum)
		fmt.Println("sha256: ", pack.SHA256)
		file, _ := os.Create("broadcast_dir/" + pack.Filename)
		defer file.Close()
		var buf = make([]byte, 32*1024)
		n, err := io.CopyBuffer(file, *conn, buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
	client.dconns <- *conn
}
