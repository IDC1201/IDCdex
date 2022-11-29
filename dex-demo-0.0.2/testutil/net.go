package testutil

import (
	"net"
	"sync"
)

var (
	mtx sync.Mutex
)

func DaemonPorts() (int, int, int) {
	mtx.Lock()
	defer mtx.Unlock()

	tmPort := FreePort()
	rpcPort := FreePort()
	p2pPort := FreePort()
	return tmPort, rpcPort, p2pPort
}

func FreePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
