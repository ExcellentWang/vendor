package payment

import (
	"pay-srv-client/tf/paymentapi"
	"pay-srv-client/tf/req"
	"pay-srv-client/tf/resp"
)

type PaymentSrv interface {
	baseService
	paymentapi.PaymentSrv
}

type paymentSrvProxy struct {
	baseServiceProxy
	paymentapi.PaymentSrv
}

func (this *paymentSrvProxy) GetPayInfo(req *req.GetPayInfoReq) (r *resp.GetPayInfoResp, err error) {
	defer this.ifAutoClose()
	//this.signHeader(req.Header)
	return this.PaymentSrv.GetPayInfo(req)
}
func (this *paymentSrvProxy) OrderQuery(req *req.OrderQueryReq) (r *resp.OrderQueryResp, err error) {
	defer this.ifAutoClose()
	//this.signHeader(req.Header)
	return this.PaymentSrv.OrderQuery(req)
}

func (this *paymentSrvProxy) CloseOrder(req *req.CloseOrderReq) (r *resp.CloseOrderResp, err error) {
	defer this.ifAutoClose()
	//this.signHeader(req.Header)
	return this.PaymentSrv.CloseOrder(req)
}
