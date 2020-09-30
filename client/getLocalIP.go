package client

import (
    "net"
)

func getLocalIP() (string, error) {
    ifaces, err := net.InterfaceAddrs()
    if err != nil {
        return  "", err
    }

    return ifaces[1].String(), nil
}