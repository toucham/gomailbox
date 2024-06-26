package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
)

var IP = os.Getenv("GOMAILBOX_IP")
var PORT = os.Getenv("GOMAILBOX_PORT")

func main() {
	// set up the program (e.g. config, global values, and etc.)
	if PORT == "" {
		PORT = "8080"
	}

	if IP == "" {
		IP = "127.0.0.1"
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc() // cancel the goroutine when the main is completed

	// setup a server to listen to requests
	ADDR := fmt.Sprintf("%s:%s", IP, PORT)
	addr, err := net.ResolveTCPAddr("tcp", ADDR)
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close() // closes the server when reaches the end

	log.Printf("Listening on %s...\n", ADDR)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf(err.Error())
		}
		go acceptConn(conn, ctx)
	}
}

func acceptConn(conn net.Conn, ctx context.Context) {
	defer conn.Close()
	select { // don't accept connection when context is cancelled
	case <-ctx.Done():
		log.Println(ctx.Err().Error())
		return
	default:
		switch conn.(type) {
		case *net.TCPConn:
			getTCPInfo(conn.(*net.TCPConn))
		default:
			log.Printf("Unsupported connection type: %T", conn)
			return
		}
	}
}

// create a client conn to send request
func getTCPInfo(conn *net.TCPConn) {
	remoteAddr := conn.RemoteAddr().String()
	localAddr := conn.LocalAddr().String()
	log.Printf("Remote address: %s", remoteAddr)
	log.Printf("Local address: %s", localAddr)
}
