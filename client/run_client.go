package client

import (
    "net"
    "log"
)

type Client struct {
    addr *string
    ip, port *string
    dir *string
    files *[]net.Conn
    aconns *map[net.Conn]bool
    dconns *chan net.Conn
    TCPconns *chan net.Conn
}

func (client *Client) Run() {
    *client.ip, _ = getLocalIP()
    *client.addr = net.JoinHostPort(*client.ip, *client.port)
    *client.dconns = make(chan net.Conn, 1)
    *client.TCPconns = make(chan net.Conn, 1)
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

