package client

import (
    "encoding/json"
    "net"
    "bufio"
    "os"
    "fmt"
)

func (client *Client) reciveConn(conn *net.Conn) {
    rd := bufio.NewReader(*conn)
    for {
        var (
            buffer = make([]byte, client.buff)
            message string
            pack *packageTCP
        )
        for {
            length, err := rd.Read(buffer)
            if err != nil {
                break
            }
            message += string(buffer[:length])
        }
        err := json.Unmarshal([]byte(message), &pack)
		if err != nil {
			break
        }
        fmt.Println(pack.From.Filename)
        file, err := os.Create("broadcast_dir/"+pack.From.Filename)
        if err != nil {
            fmt.Println(err)
        }
        file.Write(pack.Body)
    }
    client.dconns <- *conn
}
