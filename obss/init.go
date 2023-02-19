package obss

import "github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"

var obsClient *obs.ObsClient

func Init() {
	var err error

	ep := "https://" + endPoint
	obsClient, err = obs.New(ak, sk, ep)
	if err != nil {
		panic(err)
	}
}
