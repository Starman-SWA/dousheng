package obss

import (
	"dousheng/pkg/consts"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

var obsClient *obs.ObsClient

func init() {
	var err error

	ep := "https://" + consts.ObsEndPoint
	obsClient, err = obs.New(consts.ObsAk, consts.ObsSk, ep)
	if err != nil {
		panic(err)
	}
}
