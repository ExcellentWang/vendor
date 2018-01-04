include "Req.thrift"
include "Resp.thrift"

service PaymentSrv{

    /**
     * 商户系统先调用该接口在微信支付服务后台生成预支付交易单，返回正确的预支付交易回话标识后再在APP里面调起支付。
     *
     */
     Resp.GetPayInfoResp GetPayInfo(1: Req.GetPayInfoReq req)

     /**
      * 提供所有微信支付订单的查询，商户可以通过该接口主动查询订单状态，完成下一步的业务逻辑
      *
      */
      Resp.OrderQueryResp OrderQuery(1: Req.OrderQueryReq req)

      /**
      * 关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
      * 注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。
      */
      Resp.CloseOrderResp CloseOrder(1: Req.CloseOrderReq req)

      /**
      * 以下情况需要调用关单接口：商户订单支付失败需要生成新单号重新发起支付，要对原订单号调用关单，避免重复支付；系统下单后，用户支付超时，系统退出不再受理，避免用户继续，请调用关单接口。
      * 注意：订单生成后不能马上调用关单接口，最短调用时间间隔为5分钟。
      */
     // Resp.DownloadbillResp Downloadbill(1: Req.DownloadbillReq req)





}