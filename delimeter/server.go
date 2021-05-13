package delimeter

import (
	"bufio"
	"io"
	"log"
	"net"
)

func Start()  {
	listener,err := net.Listen("tcp","127.0.0.1:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for  {
		con,err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer con.Close()

		reader := bufio.NewReader(con)
		for {
			data,err := reader.ReadSlice('\n')
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}else {
					break
				}
			}
			log.Println("received msg",len(data),"bytes:",string(data))
		}
	}
}
