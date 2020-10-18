package client

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func (client *Client) reciveConn(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	sumb, _ := reader.ReadString('\n')
	sumb = strings.TrimSpace(sumb)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	if checkSum := client.Files[line]; !bytes.Equal(checkSum[:], []byte(sumb)) {
		file, _ := os.Create("broadcast_dir/" + line)
		defer file.Close()
		_, err := io.Copy(file, *conn)
		if err != nil {
			fmt.Println(err)
		}
	}
	client.dconns <- *conn
}
