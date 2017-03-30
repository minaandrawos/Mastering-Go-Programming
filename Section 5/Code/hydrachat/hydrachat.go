package hydrachat

import (
	"Hydra/hlogger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

func Run() error {
	l, err := net.Listen("tcp", ":2100")
	r := CreateRoom("HydraChat")
	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}
	go func(l net.Listener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				logger.Println("Error accepting connection from chat client", err)
				break
			}
			// fan-out
			go handleConnection(r, conn)
		}
	}(l)

	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		<-r.Msgch
		os.Exit(0)
	}()
	return nil
}

func handleConnection(r *room, c net.Conn) {
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
