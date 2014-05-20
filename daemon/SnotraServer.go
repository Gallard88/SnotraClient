package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

const (
	uSock  = "/tmp/.Snotra.Socket"
	logDir = "./"
)

type Msg struct {
	Module    string
	Date      string
	Parameter string
	Value     string
}

/*
 * Local logger.
 * Accept Unix socket
 * Receive data whil the socket is open,
 * this data is parsed and then written to a file.
 *
 */

func ClientReceiver(c net.Conn) {
	/*
	 * Here we are connected to a specific client,
	 * we wait until data is ready, then we unmarshal it into
	 * a struct
	 */
	buf := make([]byte, 4096)
	for {
		nr, err := c.Read(buf)
		if err != nil {
			c.Close()
			return
		}
		data := buf[0:nr]

		var m Msg
		err = json.Unmarshal(data, &m)
		if err != nil {
			continue
		}
		line := fmt.Sprintf("%s, %s, %s, %s\r\n", m.Module, m.Date, m.Parameter, m.Value)

		f, err := os.OpenFile(logDir+m.Module+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err == nil {
			f.Write([]byte(line))
			f.Close()
		}
	}
}

func main() {
	fmt.Printf("Snotra Online\n")

	l, err := net.Listen("unixpacket", uSock)
	if err != nil {
		println("listen error", err.Error())
		return
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			println("accept error", err.Error())
			return
		}
		go ClientReceiver(fd)
	}
}
