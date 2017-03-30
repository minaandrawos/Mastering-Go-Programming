package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
	//"Hydra/HydraConfigurator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting hydraChatClient....")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)

	/*confStruct := struct{
		Name string `name:"name"`
		RemoteAddr string `name:"remoteip"`
		TCP bool `name:"tcp"`
	}{}

	HydraConfigurator.GetConfiguration(HydraConfigurator.CUSTOM,&confStruct,"chat.conf")
	name = confStruct.Name
		proto := "tcp"
	if !confStruct.TCP{
		proto = "udp"
	}
	*/

	fmt.Printf("Hello %s, connecting to the hydra chat system.... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2100")
	if err != nil {
		log.Fatal("Could not connect to hydra chat system", err)
	}
	fmt.Println("Connected to hydra chat system")
	name += ":"
	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	/*
		for err==nil {
			msg := ""
			fmt.Print(name)
			fmt.Scan(&msg)
			msg = name+msg+"\n"
			fmt.Println("Duplicate: " + msg)
			_,err = fmt.Fprintf(conn,msg)

		}
	*/

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name+msg+"\n")
	}
}
