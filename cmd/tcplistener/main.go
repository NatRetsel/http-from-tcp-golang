package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/natretsel/http-from-tcp-golang/internal/request"
)

const port = ":42069"

func main() {

	netListener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("could not listen on port %v: %v", port, err)
		os.Exit(1)
	}
	defer netListener.Close()
	fmt.Println("Listening for TCP traffic on", port)
	for {
		netConnection, err := netListener.Accept()

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Connection accepted from ", netConnection.RemoteAddr())
		reqObj, err := request.RequestFromReader(netConnection)
		if err != nil {
			log.Fatal(err)
		}

		// print
		fmt.Println("Request line:")
		fmt.Printf("- Method: %v\n", reqObj.RequestLine.Method)
		fmt.Printf("- Target: %v\n", reqObj.RequestLine.RequestTarget)
		fmt.Printf("- Version: %v\n", reqObj.RequestLine.HttpVersion)
		fmt.Println("Connection to ", netConnection.RemoteAddr(), "closed")
	}
}
