/**
 * 请求头信息
 */
struct ReqHeader{
	/**
	 * App类型（1安卓，2ios，3H5，4网站）
	 */
	1: i32 appType
	/**
	 * App版本号
	 */
	2: string appVersion
	/**
	 * Token
	 */
	3: string token
	/**
	 * 当前时间戳(毫秒数)
	 */
	4: i64 timestamp
	/**
	 * 长度大于10位的随机字符串不能重复,推荐使用UUID
	 */
	5: string nonce
	/**
	 * 签名信息
	 */
	6: string signature
	/**
	 * 用户id
	 */
	7: i64 userId
	/**
	 * APP种类(1:非凡医生端 2:非凡患者端)
	 */
	8: i32 appKind
	/**
	 * app标识
	 */
	9: string appMark
		/**
	 * 设备Id
	 */
	10: string deviceId
	/**
	 * 设备名称
	 */
	11: string deviceName
	/**
	 * 系统名称
	 */
	12: string osName
}