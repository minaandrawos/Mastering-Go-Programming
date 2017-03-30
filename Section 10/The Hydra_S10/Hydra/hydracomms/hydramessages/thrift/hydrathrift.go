package thriftmsgs

import (
	"Hydra/hydracomms/hydramessages/thrift/gen-go/hydraThrift"
	"errors"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var thriftProtocolFactory thrift.TProtocolFactory
var thriftTransportFactory thrift.TTransportFactory

//assume some defaults
func init() {
	thriftProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	thriftTransportFactory = thrift.NewTTransportFactory()
}

type HydraShipHandler struct {
	resultChan chan interface{}
}

func NewHydraShipHandler(o chan interface{}) *HydraShipHandler {
	return &HydraShipHandler{
		resultChan: o,
	}
}

func (handler *HydraShipHandler) AddShip(s *hydraThrift.Ship) (err error) {
	handler.resultChan <- s
	return
}

func StartThriftServer(addr string, h *HydraShipHandler) (err error) {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return
	}
	processor := hydraThrift.NewHydraThriftServiceProcessor(h)
	server := thrift.NewTSimpleServer4(processor, transport, thriftTransportFactory, thriftProtocolFactory)
	return server.Serve()
}

func RunThriftClient(obj interface{}, addr string) (err error) {
	s, ok := obj.(*hydraThrift.Ship)
	if !ok {
		return errors.New("Unknown type.. ")
	}
	sock, err := thrift.NewTSocket(addr)
	if err != nil {
		return
	}
	log.Println("Thrift socket created")
	transport := thriftTransportFactory.GetTransport(sock)
	err = transport.Open()
	if err != nil {
		return
	}
	defer transport.Close()
	log.Println("Thrift transport opened")
	client := hydraThrift.NewHydraThriftServiceClientFactory(transport, thriftProtocolFactory)
	log.Println("Calling add")
	err = client.AddShip(s)
	return
}
