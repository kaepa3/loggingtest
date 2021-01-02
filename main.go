package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

var sc = bufio.NewScanner(os.Stdin)

var format = logging.MustStringFormatter(
	`time:%{time:2006/01/02-15:04:05.000}	func:%{shortfunc}	level:%{level:.4s}	msg:%{message}`,
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
		flg = process(flg)
	}
}
func process(flg bool) bool {
	if sc.Scan() {
		t := sc.Text()

		switch t {
		case "e":
			return false
		default:
			if count, e := strconv.Atoi(t); e == nil {
				for i := 0; i < count; i++ {
					if text, err := RandomString(rand.Intn(8) + 4); err == nil {
						fmt.Println(text)
						switch rand.Intn(3) {
						case 0:
							log.Error(text)
						case 1:
							log.Debug(text)
						case 2:
							log.Info(text)
						default:
							log.Warning(text)
						}
					}
				}
			} else {
				if len(t) < 10 {
					log.Debug(t)
				} else {
					log.Info(t)
				}
			}
		}
	}
	return true
}
func RandomString(digit int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

func OpenFile(path string) *os.File {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	return fp
}
