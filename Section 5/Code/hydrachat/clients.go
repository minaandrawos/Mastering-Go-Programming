package hydrachat

import (
	"bufio"
	"net"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

func StartClient(msgCh chan <- string, cn net.Conn, quit chan struct{}) (chan <- string , <- chan struct{}){
	c := new(client)
	c.Reader = bufio.NewReader(cn)
	c.Writer = bufio.NewWriter(cn)
	c.wc = make(chan string)
	done := make(chan struct{})

	//setup the reader
	go func() {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			logger.Println(scanner.Text())
			//fan-in
			msgCh <- scanner.Text()
		}
		done <- struct{}{}
	}()

	//setup the writer
	c.writeMonitor()

	go func(){
		<- quit
		cn.Close()
	}()

	return c.wc,done
}

func (c *client) writeMonitor() {
	go func() {
		for s := range c.wc {
			//logger.Println("Sending",s)
			c.WriteString(s + "\n")
			c.Flush()
		}
	}()
}
