package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandBytes generates a random string, which is then returned to the connecting
// SSH client until such time as they give up and disconnect.
func RandBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	listener, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			time.Sleep(2000 * time.Millisecond)
			for {
				fmt.Fprintf(c, "%x\r\n", RandBytes(4))
				time.Sleep(10000 * time.Millisecond)
			}
		}(conn)
	}
}
