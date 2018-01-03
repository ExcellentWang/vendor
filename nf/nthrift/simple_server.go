package nthrift

import ()

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	log "github.com/cihub/seelog"
	"runtime/debug"
)

type TSimpleServer struct {
	quit chan struct{}

	processorFactory       thrift.TProcessorFactory
	serverTransport        thrift.TServerTransport
	inputTransportFactory  thrift.TTransportFactory
	outputTransportFactory thrift.TTransportFactory
	inputProtocolFactory   thrift.TProtocolFactory
	outputProtocolFactory  thrift.TProtocolFactory
}

func NewTSimpleServer(processor thrift.TProcessor, serverTransport thrift.TServerTransport, transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory) *TSimpleServer {
	return &TSimpleServer{
		processorFactory:       thrift.NewTProcessorFactory(processor),
		serverTransport:        serverTransport,
		inputTransportFactory:  transportFactory,
		outputTransportFactory: transportFactory,
		inputProtocolFactory:   protocolFactory,
		outputProtocolFactory:  protocolFactory,
		quit: make(chan struct{}, 1),
	}
}

func (p *TSimpleServer) Serve() error {
	err := p.serverTransport.Listen()
	if err != nil {
		return err
	}
	p.acceptLoop()
	return nil
}

func (p *TSimpleServer) Stop() error {
	p.quit <- struct{}{}
	p.serverTransport.Interrupt()
	return nil
}

func (p *TSimpleServer) acceptLoop() error {
	for {
		client, err := p.serverTransport.Accept()
		if err != nil {
			select {
			case <-p.quit:
				return nil
			default:
			}
			return err
		}
		if client != nil {
			go func() {
				if err := p.processRequests(client); err != nil {
					log.Error("error processing request:", err)
				}
			}()
		}
	}
}

func (p *TSimpleServer) processRequests(client thrift.TTransport) error {
	processor := p.processorFactory.GetProcessor(client)
	inputTransport := p.inputTransportFactory.GetTransport(client)
	outputTransport := p.outputTransportFactory.GetTransport(client)
	inputProtocol := p.inputProtocolFactory.GetProtocol(inputTransport)
	outputProtocol := p.outputProtocolFactory.GetProtocol(outputTransport)
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("panic in processor: %s: %s", e, debug.Stack())
		}
	}()
	if inputTransport != nil {
		defer inputTransport.Close()
	}
	if outputTransport != nil {
		defer outputTransport.Close()
	}
	for {
		ok, err := processor.Process(inputProtocol, outputProtocol)
		if err, ok := err.(thrift.TTransportException); ok && err.TypeId() == thrift.END_OF_FILE {
			return nil
		} else if err != nil {
			log.Errorf("error processing request: %s", err)
			return err
		}
		if !ok {
			break
		}
	}
	return nil
}
