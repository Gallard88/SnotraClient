package main

import (
	"encoding/json"
	"fmt"
	"net"
)

const (
	uSock  = "/tmp/.Snotra.Socket"
	LogDir = "./"
)

/********************************
 *
 * Local logger.
 * Accept Unix socket
 * Receive data whil the socket is open,
 * this data is parsed and then written to a file.
 *
 */

func ClientReceiver(c net.Conn, packets chan ClientMsg) {
	/*
	 * Here we are connected to a specific client,
	 * we wait until data is ready, then we unmarshal it into
	 * a struct, and insert it into the chanel.
	 */
	buf := make([]byte, 4096)
	for {
		nr, err := c.Read(buf)
		if err != nil {
			c.Close()
			return
		}
		data := buf[0:nr]

		var m ClientMsg
		err = json.Unmarshal(data, &m)
		if err != nil {
			continue
		}
		packets <- m
	}
}

func main() {
	fmt.Printf("Snotra Online\n")
  MsgMux_SetLoggingDir(LogDir)

	l, err := net.Listen("unixpacket", uSock)
	if err != nil {
		println("listen error", err.Error())
		return
	}

	// Daemonise here....

	incoming := make(chan ClientMsg, 32)
	go MessageHandler(incoming)

	for {
		fd, err := l.Accept()
		if err != nil {
			println("accept error", err.Error())
			return
		}
		go ClientReceiver(fd, incoming)
	}
}
