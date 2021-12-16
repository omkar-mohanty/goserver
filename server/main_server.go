package server

import (
	"net"
	"sync"
)

var (
	tcpConnections []net.Conn
	lock           sync.Mutex
)

func StartServer(port string) {
	tcpAddr := tcpAddrResolver(port)
	setupListner(tcpAddr)
}
func handleConn(conn net.Conn) {
	var buf [512]byte
	for {

		n, err := conn.Read(buf[:])
		checkError(err)
		go writeToAllConn(buf[0:n])
	}
}
func writeToAllConn(buf []byte) {
	for _, conn := range tcpConnections {
		go conn.Write(buf)
	}
}
func acceptConns(listner *net.TCPListener) {
	for {
		conn, err := listner.Accept()
		tcpConnections = append(tcpConnections, conn)
		addLogEntry("New Connection added " + conn.LocalAddr().String())
		checkError(err)
		go handleConn(conn)
	}
}
func setupListner(tcpAddr *net.TCPAddr) {
	listner, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	acceptConns(listner)
}
func tcpAddrResolver(port string) *net.TCPAddr {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	checkError(err)
	return tcpAddr
}

func checkError(err error) {
	if err != nil {
		addLogEntry("[-]Fatal Error: " + err.Error())
		return
	}
}
