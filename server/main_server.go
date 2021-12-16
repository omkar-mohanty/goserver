package server

import (
	"C/Users/KIIT/OpenSource/gitrepos/GoServer/admin"
	"fmt"
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
		if err != nil {
			s := fmt.Sprintf("[*]Read message from:%s , message:%s",
				conn.RemoteAddr().String(),
				buf[:0])
			admin.Log(s)
			go writeToAllConn(buf[:n])
		}
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
		if err != nil {
			s := fmt.Sprintf("[+]New Connection:%s ",
				conn.RemoteAddr().String())
			admin.Log(s)
			go handleConn(conn)
		}
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
		return
	}
}
