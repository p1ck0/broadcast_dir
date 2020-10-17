package client

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"os"
	"strings"
)

func (client *Client) reciveConn(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	sum, _ := reader.ReadString('\n')
	sum = strings.TrimSpace(sum)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	if checkSum := client.files[line]; !bytes.Equal(checkSum[:], []byte(sum)) {
		file, _ := os.Create("broadcast_dir/" + line)
		io.Copy(file, *conn)
	}
	client.dconns <- *conn
}
