package client

import (
    "encoding/json"
    "net"
    "bufio"
    "log"
    "os"
    "fmt"
)

func (client *Client) reciveConn(conn *net.Conn) {
    rd := bufio.NewReader(*conn)
    for {
        var (
            buffer = make([]byte, *client.buff)
            message string
            pack *packageTCP
        )
        length, err := rd.Read(buffer)
		if err != nil {
			break
        }
        message += string(buffer[:length])
        err = json.Unmarshal([]byte(message), &pack)
		if err != nil {
			log.Println(err)
        }
        switch {
        case !*pack.From.Mod:
            (*client.aconns)[*conn] = false
        case *pack.From.Mod:
            fmt.Println(*(pack).From.Filename)
            file, _ := os.Create("broadcast_dir/"+(*(pack).From.Filename))
            file.Write(*pack.Body)
        }
    } 
}