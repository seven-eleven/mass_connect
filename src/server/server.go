package main

import (
	"common"
	"net"
	"os"
)

func main() {
	// get args
	args := os.Args
	argc := len(os.Args)
	if 3 != argc {
		common.Log("Args num for server error.")
		return
	}

	// tcp listen
	ip := args[1]
	port := args[2]
	server := ip + ":" + port
	netListen, err := net.Listen("tcp", server)
	if nil != err {
		common.Log("Tcp listen error.")
		return
	}

	defer netListen.Close()

	// listen and interact
	var conn net.Conn
	for {
		common.Log("Waiting for clients")
		conn, err = netListen.Accept()
		if nil != err {
			continue
		}

		common.Log(conn.RemoteAddr().String(), " tcp connect success.")

		go serverHandle(conn)
	}
}

func serverHandle(conn net.Conn) {
	if nil == conn {
		return
	}

	for {
		err := common.RecvMsg(conn)
		if nil != err {
			return
		}

		ack := "msg ack."
		err = common.SendMsg(conn, ack)
		if nil != err {
			return
		}
	}

	common.Log("Server handle thread over.")
}
