package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8000", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runClient(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What message would you like to send?")

	for scanner.Scan() && err == nil {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(append(scanner.Bytes(), '\r'))

		fmt.Println("What message would you like to send?")
		buffer := make([]byte, 1024)
		//conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buffer))
	}
	return err
}

func runServer(address string) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log.Println("Listening..... ")
	defer l.Close()
	for {
		c, err := l.Accept()
		if err!= nil {
			return err
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)
	for {
		//buffer := make([]byte, 1024)
		c.SetDeadline(time.Now().Add(5 * time.Second))
		line, err := reader.ReadString('\r')
		//_, err := c.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Received %s from address %s \n", line[:len(line)-1], c.RemoteAddr())
		writer.WriteString("Message received... ")
		writer.Flush()
	}
}
