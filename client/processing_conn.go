package client

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func (client *Client) reciveConn(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	fmt.Println(line)
	file, _ := os.Create("broadcast_dir/" + line)
	io.Copy(file, *conn)
	client.dconns <- *conn
}
