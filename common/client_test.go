package common

import (
	"log"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestTcp(t *testing.T)  {
	conn,err := net.Dial("tcp","127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for i := 0;i < 10;i++ {
		var err error
		_, err = conn.Write([]byte(strconv.Itoa(i) + "aaa\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + "bbbb\n"))
		_, err = conn.Write([]byte(strconv.Itoa(i) + "ccccc\n"))
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second)
}
