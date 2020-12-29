package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

var sc = bufio.NewScanner(os.Stdin)

var format = logging.MustStringFormatter(
	`%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{message}`,
)

func main() {
	fmt.Println("start")
	f := OpenFile("log.log")
	defer f.Close()
	backend := logging.NewLogBackend(f, "", 0)
	b := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(b)

	flg := true
	for flg {
		if sc.Scan() {
			t := sc.Text()
			switch t {
			case "e":
				flg = false
			default:
				if len(t) < 10 {
					log.Debug(t)
				} else {
					log.Info(t)
				}
			}
		} else {
			break
		}
	}
}

func OpenFile(path string) *os.File {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	return fp
}
