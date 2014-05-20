package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

const (
	uSock = "/tmp/.Snotra.Socket"
)

type Msg struct {
	Module string
	Date   string
  Parameter  string
	Value  string
}

func main() {

	argv := os.Args

	if len(argv) < 4 {
		fmt.Printf("Usage: SnotraClient module parameter value\n")
		return
	}

	conn, err := net.Dial("unixpacket", uSock)
	if err != nil {
		panic(err.Error())
	}

	var m Msg
	m.Module = argv[1]
	m.Parameter = argv[2]
	m.Value = argv[3]
	m.Date = time.Now().Format("2006-01-02, 15:04:05")

	data, err := json.Marshal(m)
	if err != nil {
		println(err.Error())
		return
	}

	_, err = conn.Write(data)
	if err != nil {
		println(err.Error())
		return
	}
	conn.Close()
}
