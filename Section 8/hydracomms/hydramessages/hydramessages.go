package hydramessages

import (
	"Hydra/hydracomms/hydramessages/protobuff"
	"Hydra/hydracomms/hydramessages/thrift"
	"bufio"
	"io/ioutil"
	"log"
	"net"
	"Hydra/hydracomms/hydramessages/gob"
)

//Communication messages types
const (
	Protobuf uint8 = iota
	GOB
	THRIFT
)

func EncodeAndSend(serType uint8, obj interface{}, destination string) (err error) {
	var buffer []byte
	switch serType {
	case Protobuf:
		buffer, err = protobuff.EncodeProto(obj)
		if err != nil {
			return
		}
		err = sendmessage(buffer, destination)
	case GOB:
		err = sendGobMessage(obj,destination)
	case THRIFT:
		err = thriftmsgs.RunThriftClient(obj, destination)
	}

	return
}

func sendmessage(buffer []byte, destination string) error {
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("Sending %d bytes to %s \n", len(buffer), destination)
	_, err = conn.Write(buffer)
	return err
}

func sendGobMessage(obj interface{}, destination string)error{
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer conn.Close()
	return hydragob.EncodeAndWriteGob(obj,conn)
}


func ListenAndDecode(serType uint8, listenaddress string) chan interface{} {
	out := make(chan interface{})
	if serType == THRIFT {
		go func() {
			handler := thriftmsgs.NewHydraShipHandler(out)
			if err := thriftmsgs.StartThriftServer(listenaddress, handler); err != nil {
				log.Println(err)
				close(out)
			}
		}()
		log.Println("Returning out... ")
		return out
	}

	l, _ := net.Listen("tcp", listenaddress)
	log.Println("Listening... ")

	go func() {
		defer l.Close()
		for {
			c, _ := l.Accept()

			go func(c net.Conn) {
				defer c.Close()
				r := bufio.NewReader(c)
				for {
					var err error
					var obj interface{}
					buffer, err := ioutil.ReadAll(r)

					if err != nil {
						break
					} else if len(buffer) == 0 {
						continue
					}
					switch serType {
					case Protobuf:
						log.Println("Decoding protobuf")
						obj, err = protobuff.DecodeProto(buffer)
					case GOB:
						obj, err = hydragob.DecodeAndReadGob(buffer)
					}
					if err == nil {
						out <- obj
					}
				}
			}(c)
		}
	}()

	return out
}
