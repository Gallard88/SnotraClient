package main

import (
  "fmt"
  "os"
)

type ClientMsg struct {
	Module    string
	Date      string
	Parameter string
	Value     float32
}

var (
  logDir string
)

func MsgMux_SetLoggingDir(dir string) {
  logDir = dir
}

func MessageHandler(packets chan ClientMsg) {
// Here we grab all incoming messages and sort them.
// This is a multiplexer of sorts.

	for {

		m := <-packets
    filename := logDir+m.Module+"."+m.Parameter+".log"
		line := fmt.Sprintf("%s, %f\r\n", m.Date, m.Value)
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err == nil {
			f.Write([]byte(line))
			f.Close()
		}
	}
}
