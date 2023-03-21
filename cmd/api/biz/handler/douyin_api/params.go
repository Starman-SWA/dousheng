package douyin_api

import "mime/multipart"

type PublishActionFilePara struct {
	F *multipart.FileHeader `form:"data"`
}
