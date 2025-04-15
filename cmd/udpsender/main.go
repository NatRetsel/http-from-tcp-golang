package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const networkAddr = "localhost:42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", networkAddr)
	if err != nil {
		log.Fatalf("error resolving UDP address %v: %v\n", networkAddr, err)
		os.Exit(1)
	}
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("error dialing UDP: %v\n", err)
		os.Exit(1)
	}
	defer udpConn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		inStr, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("error reading from input: %v\n", err)
			os.Exit(1)
		}
		_, err = udpConn.Write([]byte(inStr))
		if err != nil {
			log.Fatalf("error writing to UDP connection: %v\n", err)
			os.Exit(1)
		}
	}

}
