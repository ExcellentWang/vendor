package fdfs

import (
	"fdfs-srv-client/tf/proto"
	"fmt"
	"time"
)

var (
	addr    string
	signKey string
)

func InitFdfs(srvAddr string, srvKey string) {
	addr = srvAddr
	signKey = srvKey
	//addr = "192.168.1.151:15261"
	//signKey = "fdfsSrv"
}

func NewFdfsSrv(clientTimeout time.Duration, autoClose bool) (FdfsSrv, error) {
	baseProxy, err := createBaseServiceProxy("FdfsSrv", addr, clientTimeout, autoClose)
	if err != nil {
		fmt.Errorf("error is %s", err)
		return nil, err
	}
	//NewFdfsSrvClientProtocol
	//NewFdfsServiceClientProtocol
	srv := proto.NewFdfsSrvClientProtocol(baseProxy.transport, baseProxy.protocol, baseProxy.protocol)
	return &fdfsSrvProxy{baseServiceProxy: baseProxy, FdfsSrv: srv}, nil
}
