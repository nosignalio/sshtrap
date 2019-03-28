package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// RandBytes generates a random string, which is then returned to the connecting
// SSH client until such time as they give up and disconnect.
func RandBytes(n int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// Logging
	Formatter := new(log.TextFormatter)
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	// Configuration management
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$HOME/.sshtrap")
	viper.SetConfigName("sshtrap")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file ", err)
	}

	ln, err := net.Listen(viper.GetString("service.protocol"), viper.GetString("service.port"))
	if err != nil {
		log.Panic(err)
	}

	log.Info("Listening for connections on ", ln.Addr())
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	log.Info("Connection from ", conn.RemoteAddr())

	defer conn.Close()

	for {
		fmt.Fprintf(conn, "%x\r\n", RandBytes(4))
		time.Sleep(10000 * time.Millisecond)
	}
}
