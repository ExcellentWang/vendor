package fdfs

import (
	"fdfs-srv-client/tf/proto"
)

type FdfsSrv interface {
	baseService
	proto.FdfsSrv
}

type fdfsSrvProxy struct {
	baseServiceProxy
	proto.FdfsSrv
}

func (this *fdfsSrvProxy) Upload(req *proto.UploadReq) (r *proto.UploadResp, err error) {
	defer this.ifAutoClose()
	this.signHeader(req.Header)
	return this.FdfsSrv.Upload(req)
}
func (this *fdfsSrvProxy) Download(req *proto.DownloadReq) (r *proto.DownloadResp, err error) {
	defer this.ifAutoClose()
	this.signHeader(req.Header)
	return this.FdfsSrv.Download(req)
}

func (this *fdfsSrvProxy) DeleteFile(req *proto.DeleteReq) (r *proto.DeleteResp, err error) {
	defer this.ifAutoClose()
	this.signHeader(req.Header)
	return this.FdfsSrv.DeleteFile(req)
}
