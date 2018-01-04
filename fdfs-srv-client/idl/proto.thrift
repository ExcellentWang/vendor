namespace java pingan.fdfsagent.tf

const string VERSION = "1.0.0"

struct ReqHeader {
	1: i64 timestamp
	/**
	 * 长度大于10位的随机字符串不能重复,推荐使用UUID
	 */
	2: string traceId
	/**
	 * 签名信息 Sha256Hex(sigKey + timestamp + nonce)
	 */
	3: string sig
}

struct ResqHeader {
}

struct UploadReq {
	1: ReqHeader header
	/**
	 * 当前支持两个(viodoc:公开文件组可以通过nginx访问文件, private:私有文件组,上传后必须通过下载接口才可访问)
	 */
	2: string groupName
	/**
	 * 文件内容
	 */
	3: binary fileContent
	/**
	 * 文件扩展名如png,jpg等文件的扩展名
	 */
	4: string fileExtName
	/**
	 * 主文件名,如果为slave上传时需要指定
	 */
	5: string masterFilename
	/**
	 * 从文件前缀名如: _100x100,如果为slave上传时需要指定
	 */
	6: string slavePrefixName
	/**
	 * time to live 存活时间（天）
	 */
	7: i32 ttl = -1
	/**
	 * 压缩后宽度,只有png或jpg后缀的图片才会被压缩,并且resizeWidth和resizeHeight二者中有一个大于0才会被压缩
	 */
	8: i32 resizeWidth = 0
	/**
	 * 压缩后高度,只有png或jpg后缀的图片才会被压缩,并且resizeWidth和resizeHeight二者中有一个大于0才会被压缩
	 */
	9: i32 resizeHeight = 0
}

struct UploadResp {
	1: ResqHeader header
	2: string groupName
	3: string remoteFilename
}


struct DownloadReq {
	1: ReqHeader header
	2: string groupName
	3: string remoteFilename
}

struct DownloadResp {
	1: ResqHeader header
	2: binary fileContent
	3: i64 downloadSize
}

struct DeleteReq {
	1: ReqHeader header
	2: string groupName
	3: string remoteFilename

}

struct DeleteResp {
	1: ResqHeader header
}


service FdfsSrv {
	
	UploadResp upload(1: UploadReq req)
	
	DownloadResp download(1: DownloadReq req)

	DeleteResp deleteFile(1: DeleteReq req)
}