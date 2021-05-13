package length

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

func Start() {
	listener, err := net.Listen("tcp", "0.0.0.0:8866")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		defer con.Close()

		reader := bufio.NewReader(con)
		for {
			peek,err := reader.Peek(4)
			if err != nil {
				if err != io.EOF {
					log.Println(err)
				}else {
					break
				}
			}

			buffer := bytes.NewBuffer(peek)
			var length int32
			err = binary.Read(buffer,binary.BigEndian,&length)
			if err != nil {
				log.Println(err)
			}
			if int32(reader.Buffered()) < length + 4 {
				continue
			}

			data := make([]byte, length + 4)
			_, err = reader.Read(data)
			if err != nil {
				continue
			}
			log.Println("received msg", string(data[4:]))
		}
	}
}
