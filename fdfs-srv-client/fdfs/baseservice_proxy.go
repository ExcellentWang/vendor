package fdfs

import (
	"fdfs-srv-client/tf/proto"
	"fmt"
	"io"
	"nf/util/crypto"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type baseService interface {
	io.Closer
}

type baseServiceProxy struct {
	autoClose bool
	transport thrift.TTransport
	protocol  thrift.TProtocol
}

func (srv *baseServiceProxy) Close() error {
	return srv.transport.Close()
}

func (srv *baseServiceProxy) ifAutoClose() {
	if srv.autoClose {
		srv.transport.Close()
	}
}

func (srv *baseServiceProxy) signHeader(header *proto.ReqHeader) {
	header.Timestamp = time.Now().UnixNano() / int64(time.Millisecond)
	header.Sig = crypto.Sha256Hex(fmt.Sprintf("%s%d%s", signKey, header.Timestamp, header.TraceId))
}

func createBaseServiceProxy(serviceName string, addr string, clientTimeout time.Duration, autoClose bool) (proxy baseServiceProxy, err error) {
	var transport thrift.TTransport
	transport, err = thrift.NewTSocketTimeout(addr, clientTimeout)
	if err != nil {
		return proxy, err
	}
	// protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	protocol := protocolFactory.GetProtocol(transport)
	err = transport.Open()
	if err != nil {
		return proxy, err
	}
	return baseServiceProxy{
		autoClose: autoClose,
		transport: transport,
		protocol:  protocol}, nil
}
