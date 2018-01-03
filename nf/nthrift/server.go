package nthrift

import (
	"crypto/tls"
	"git.apache.org/thrift.git/lib/go/thrift"
	"time"
)

type Server interface {
	Serve(addr string, processor thrift.TProcessor) error
	ServeTLS(addr string, processor thrift.TProcessor, serverCrt []byte, serverKey []byte) error
}

func NewServer(protocolFactory thrift.TProtocolFactory, transportFactory thrift.TTransportFactory, clientTimeout time.Duration) Server {
	return &defaultServer{
		protocolFactory:  protocolFactory,
		transportFactory: transportFactory,
		clientTimeout:    clientTimeout}
}

type defaultServer struct {
	protocolFactory  thrift.TProtocolFactory
	transportFactory thrift.TTransportFactory
	clientTimeout    time.Duration
}

func (this *defaultServer) Serve(addr string, processor thrift.TProcessor) error {
	transport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	return this.serve(transport, processor)
}

func (this *defaultServer) ServeTLS(addr string, processor thrift.TProcessor, serverCrt []byte, serverKey []byte) error {
	cfg := new(tls.Config)
	cert, err := tls.X509KeyPair(serverCrt, serverKey)
	if err != nil {
		return err
	}
	cfg.Certificates = append(cfg.Certificates, cert)
	transport, err := thrift.NewTSSLServerSocket(addr, cfg)
	if err != nil {
		return err
	}
	server := NewTSimpleServer(processor, transport, this.transportFactory, this.protocolFactory)
	return server.Serve()
}

func (this *defaultServer) serve(transport thrift.TServerTransport, processor thrift.TProcessor) error {
	server := NewTSimpleServer(processor, transport, this.transportFactory, this.protocolFactory)
	return server.Serve()
}
