package client

import (
	"fmt"
    "net"
    "log"
)

//Client is the structure of the client
type Client struct {
    addr string
    ip, Port string
    dir *string
    buff int
    files *[]net.Conn
    dconns chan net.Conn
    TCPconns chan net.Conn
}


//Run func for run a client on a local network
func (client *Client) Run() {
    client.ip = "127.0.0.1"
    fmt.Println(client.ip)
    client.addr = net.JoinHostPort(client.ip, client.Port)
    client.dconns = make(chan net.Conn, 1)
    client.TCPconns = make(chan net.Conn, 1)
    client.buff = 1024
    client.listen()
}

func (client *Client) listen() {
    tcpaddr, _ := net.ResolveTCPAddr("tcp", client.addr)
    ln, _ := net.ListenTCP("tcp", tcpaddr)
    defer ln.Close()
    go func() {
        for {
            conn, err := ln.Accept()
			if err != nil {
				log.Fatalln(err.Error())
			}
			client.TCPconns <- conn
        }
    }()
    client.handle()
}

func (client *Client) handle() {
    for {
        select {
        case conn := <-client.TCPconns:
            go client.reciveConn(&conn)
        case dconn := <-client.dconns:
            dconn.Close()
        }
    }
}
