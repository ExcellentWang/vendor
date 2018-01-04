
struct GetPayInfoResp{
       /**
    * 返回业务层状态码
    **/
    1:i32 status
    /**
    * 返回消息
    **/
    2:string msg

    /**
    *以下字段在return_code为SUCCESS的时候有返回
    */
    /**
     * 字段名:应用APPID
     * 描述:调用接口提交的应用ID
     * 示例:wx8888888888888888
     * 必须:是
     */
    3:string appid
    /**
     * 字段名:商户号
     * 描述:调用接口提交的商户号
     * 示例:1900000109
     * 必须:是
     */
    4:string mchID
    /**
     * 字段名:包名
     */
    5:string packageName
    /**
     * 字段名:随机字符串
     * 描述:微信返回的随机字符串
     * 示例:5K8264ILTKCH16CQ2502SI8ZNMTM67VS
     * 必须:是
     */
    6:string nonceStr
    /**
     * 字段名:签名
     * 描述:微信返回的签名，详见签名算法
     * 示例:C380BEC2BFD727A4B6845133519F3AD6
     * 必须:是
     */
    7:string sign
    /**
     * 字段名:时间戳
     */
    8:string timeStamp
    /**
     * 字段名:预支付交易会话标识
     * 描述:微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
     * 示例:wx201410272009395522657a690389285100
     * 必须:是
     */
    9:string prepayId
}

struct OrderQueryResp{

     /**
     * 字段名:返回状态码
     * 描述:SUCCESS
     * 示例:SUCCESS/FAIL,此字段是通信标识，非交易标识，交易是否成功需要查看trade_state来判断
     * 必须:是
     */
     1:string return_code
     /**
     * 字段名:返回信息
     * 描述:   返回信息，如非空，为错误原因
     *         签名失败
     *         参数格式校验错误
     * 示例:签名失败
     * 必须:否
     */
     2:string return_msg

//    以下字段在return_code为SUCCESS的时候有返回

    /**
    * 字段名:应用APPID
    * 描述:微信开放平台审核通过的应用APPID
    * 示例:wxd678efh567hg6787
    * 必须:是
    */
    3:string appid
    /**
    * 字段名:商户号
    * 描述:1230000109
    * 示例:微信支付分配的商户号
    * 必须:是
    */
    4:string mch_id
    /**
    * 字段名:随机字符串
    * 描述:随机字符串，不长于32位。推荐随机数生成算法
    * 示例:5K8264ILTKCH16CQ2502SI8ZNMTM67VS
    * 必须:是
    */
    5:string nonce_str
    /**
    * 字段名:签名
    * 描述:签名，详见签名生成算法
    * 示例:C380BEC2BFD727A4B6845133519F3AD6
    * 必须:是
    */
    6:string sign
     /**
     * 字段名:业务结果
     * 描述:SUCCESS/FAIL
     * 示例:SUCCESS
     * 必须:是
     */
     7:string result_code
     /**
     * 字段名:错误代码
     * 描述:错误码
     * 示例:SYSTEMERROR
     * 必须:否
     */
     8:string err_code
     /**
     * 字段名:错误代码描述
     * 描述:	结果信息描述
     * 示例:	系统错误
     * 必须:否
     */
     9:string err_code_des
    //以下字段在return_code 和result_code都为SUCCESS的时候有返回
     /**
     * 字段名:设备号
     * 描述:微信支付分配的终端设备号
     * 示例:013467007045764
     * 必须:否
     */
     10:string device_info
     /**
     * 字段名:用户标识
     * 描述:用户在商户appid下的唯一标识
     * 示例:oUpF8uMuAJO_M2pxb1Q9zNjWeS6o
     * 必须:是
     */
     11:string openid
     /**
     * 字段名:是否关注公众账号
     * 描述:用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
     * 示例:Y
     * 必须:否
     */
     12:string is_subscribe
     /**
     * 字段名:交易类型
     * 描述:调用接口提交的交易类型
     * 示例:APP
     * 必须:是
     */
     13:string trade_type

    /**
    * 字段名:交易状态
    * 描述:SUCCESS—支付成功
    *     REFUND—转入退款
    *     NOTPAY—未支付
    *     CLOSED—已关闭
    *     REVOKED—已撤销（刷卡支付）
    *     USERPAYING--用户支付中
    *     PAYERROR--支付失败(其他原因，如银行返回失败)
    * 示例:SUCCESS
    * 必须:是
    */
    14:string trade_state
     /**
     * 字段名:付款银行
     * 描述:银行类型，采用字符串类型的银行标识
     * 示例:CMC
     * 必须:是
     */
     15:string bank_type
     /**
     * 字段名:总金额
     * 描述:订单总金额，单位为分
     * 示例:100
     * 必须:是
     */
     16:i64 total_fee
     /**
     * 字段名:货币种类
     * 描述:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
     * 示例:CNY
     * 必须:否
     */
     17:string fee_type
     /**
     * 字段名:现金支付金额
     * 描述:现金支付金额订单现金支付金额，详见支付金额
     * 示例:100
     * 必须:是
     */
     18:i64 cash_fee
     /**
     * 字段名:现金支付货币类型
     * 描述:货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
     * 示例:CNY
     * 必须:否
     */
     19:string cash_fee_type
     /**
     * 字段名:应结订单金额
     * 描述:当订单使用了免充值型优惠券后返回该参数，应结订单金额=订单金额-免充值优惠券金额
     * 示例:100
     * 必须:否
     */
     20:i64 settlement_total_fee
     /**
     * 字段名:代金券金额
     * 描述:“代金券或立减优惠”金额<=订单总金额，订单总金额-“代金券或立减优惠”金额=现金支付金额，详见支付金额
     * 示例:100
     * 必须:否
     */
     21:i64 coupon_fee
     /**
     * 字段名:代金券使用数量
     * 描述:代金券或立减优惠使用数量
     * 示例:1
     * 必须:否
     */
     22:i64 coupon_count
     /**
     * 字段名:代金券ID(coupon_id_$n)
     * 描述:代金券或立减优惠ID, $n为下标，从0开始编号
     * 示例:10000
     * 必须:否
     */
     23:string coupon_id_n

    /**
     * 字段名:代金券类型(coupon_type_$n)
     * 描述:CASH--充值代金券
     *        NO_CASH---非充值优惠券
     *        开通免充值券功能，并且订单使用了优惠券后有返回（取值：CASH、NO_CASH）。$n为下标,从0开始编号，举例：coupon_type_$0
     * 示例:CASH
     * 必须:否
     */
     24:string coupon_type_n
    /**
     * 字段名:单个代金券支付金额(coupon_fee_$n)
     * 描述:单个代金券或立减优惠支付金额, $n为下标，从0开始编号
     * 示例:100
     * 必须:否
     */
     25:i64 coupon_fee_n
     /**
     * 字段名:微信支付订单号
     * 描述:微信支付订单号
     * 示例:1009660380201506130728806387
     * 必须:是
     */
     26:string transaction_id
     /**
     * 字段名:商户订单号
     * 描述:商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一
     * 示例:20150806125346
     * 必须:是
     */
     27:string out_trade_no
     /**
     * 字段名:附加数据
     * 描述:附加数据，原样返回
     * 示例:深圳分店
     * 必须:否
     */
     28:string attach
      /**
     * 字段名:支付完成时间
     * 描述:订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
     * 示例:20141030133525
     * 必须:是
     */
     29:string time_end

     /**
     * 字段名:交易状态描述
     * 描述:对当前查询订单状态的描述和下一步操作的指引
     * 示例:支付失败，请重新下单支付
     * 必须:是
     */
     30:string trade_state_desc
}
    /**
    * 关闭订单接口返回参数
    **/
    struct CloseOrderResp{
    /**
     * 字段名:返回状态码
     * 描述:SUCCESS/FAIL
     * 示例:SUCCESS
     * 必须:是
     */
     1:string return_code
     /**
      * 字段名:返回信息
      * 描述:返回信息，如非空，为错误原因
      *     签名失败
      *     参数格式校验错误
      * 示例:签名失败
      * 必须:否
      */
      2:string return_msg

    //以下字段在return_code为SUCCESS的时候有返回

     /**
      * 字段名:应用ID
      * 描述:微信开放平台审核通过的应用APPID
      * 示例:wx8888888888888888
      * 必须:是
      */
      3:string appid
     /**
      * 字段名:商户号
      * 描述:微信支付分配的商户号
      * 示例:1900000109
      * 必须:是
      */
      4:string mch_id
     /**
      * 字段名:随机字符串
      * 描述:随机字符串，不长于32位
      * 示例:5K8264ILTKCH16CQ2502SI8ZNMTM67VS
      * 必须:是
      */
      5:string nonce_str
     /**
      * 字段名:签名
      * 描述:签名，验证签名算
      * 示例:C380BEC2BFD727A4B6845133519F3AD6
      * 必须:是
      */
      6:string sign
     /**
      * 字段名:错误代码
      * 描述:详细参见微信错误列表
      * 示例:SYSTEMERROR
      * 必须:否
      */
      7:string err_code
       /**
        * 字段名:错误代码描述
        * 描述:结果信息描述
        * 示例:系统错误
        * 必须:否
        */
        30:string   err_code_des
    }





















