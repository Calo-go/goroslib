package goroslib

import (
	"sync"
)

func (n *Node) runTcprosServer(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		client, err := n.tcprosServer.Accept()
		if err != nil {
			break
		}

		n.tcpClientNew <- client
	}
}
