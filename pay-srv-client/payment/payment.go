package payment

import (
	"fmt"
	"time"
	"pay-srv-client/tf/paymentapi"
)

var (
	addr    string
	signKey string
)

func InitPayment(srvAddr string, srvKey string) {
	addr = srvAddr
	signKey = srvKey
	//addr = "192.168.1.151:15261"
	//signKey = "fdfsSrv"
}

func NewPaymentSrv(clientTimeout time.Duration, autoClose bool) (PaymentSrv, error) {
	baseProxy, err := createBaseServiceProxy("PaymentSrv", addr, clientTimeout, autoClose)
	if err != nil {
		fmt.Errorf("error is %s", err)
		return nil, err
	}
	srv := paymentapi.NewPaymentSrvClientProtocol(baseProxy.transport, baseProxy.protocol, baseProxy.protocol)
	return &paymentSrvProxy{baseServiceProxy: baseProxy, PaymentSrv: srv}, nil
}
