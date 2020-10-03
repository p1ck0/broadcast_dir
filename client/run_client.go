package client

import (
	"fmt"
    "net"
    "log"
)

//Client is the structure of the client
type Client struct {
    addr *string
    ip, port *string
    dir *string
    buff *int
    files *[]net.Conn
    aconns *map[net.Conn]bool
    dconns *chan net.Conn
    TCPconns *chan net.Conn
}


//Run func for run a client on a local network
func (client *Client) Run() {
    *client.ip, _ = getLocalIP()
    *client.addr = net.JoinHostPort(*client.ip, *client.port)
    *client.dconns = make(chan net.Conn, 1)
    *client.TCPconns = make(chan net.Conn, 1)
    *client.buff = 2048
    client.listen()
}

func (client *Client) listen() {
    tcpaddr, _ := net.ResolveTCPAddr("tcp", *client.addr)
    ln, _ := net.ListenTCP("tcp", tcpaddr)
    defer ln.Close()
    go func() {
        for {
            conn, err := ln.Accept()
			if err != nil {
				log.Fatalln(err.Error())
			}
			*client.TCPconns <- conn
        }
    }()
}

func (client *Client) handle() {
    for {
        select {
        case conn := <-*client.TCPconns:
            fmt.Println(conn)
            go client.reciveConn(&conn)
        case dconn := <-*client.dconns:
            defer dconn.Close()
            log.Printf("Client %v was gone\n", dconn)
            delete(*client.aconns, dconn)
        }
    }
}
