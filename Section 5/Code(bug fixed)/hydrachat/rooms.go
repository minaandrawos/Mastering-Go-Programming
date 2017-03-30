package hydrachat

import (
	"fmt"
	"net"
	"sync"
)

type room struct {
	name    string
	Msgch   chan string
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[chan<- string]struct{}),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

func (r *room) AddClient(c net.Conn) {
	logger.Println("Adding client", c.RemoteAddr())
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	//remove client when done is signalled
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) ClCount() int {
	return len(r.clients)
}

func (r *room) RemoveClient(wc chan<- string) {
	logger.Println("Removing client ")
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()
	select {
	case <-r.Quit:
		if len(r.clients) == 0 {
			close(r.Msgch)
		}
	default:
	}
}

func (r *room) Run() {
	logger.Println("Starting chat room", r.name)
	go func() {
		for msg := range r.Msgch {
			r.broadcastMsg(msg)
		}
	}()
}

func (r *room) broadcastMsg(msg string) {
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message: ", msg)
	for wc, _ := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
