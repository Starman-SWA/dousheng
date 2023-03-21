package obss

import (
	"dousheng/pkg/consts"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func PutFile(key string, source string) error {
	input := &obs.PutFileInput{}
	input.Bucket = consts.ObsBucketName
	input.Key = key
	input.SourceFile = source
	output, err := obsClient.PutFile(input)

	if err == nil {
		fmt.Printf("RequestId:%s\n", output.RequestId)
		fmt.Printf("ETag:%s, StorageClass:%s\n", output.ETag, output.StorageClass)
	} else {
		if obsError, ok := err.(obs.ObsError); ok {
			fmt.Println(obsError.Code)
			fmt.Println(obsError.Message)
		} else {
			fmt.Println(err)
		}
	}

	return err
}

func GenGetURL(key string) string {
	return "https://" + consts.ObsBucketName + "." + consts.ObsEndPoint + "/" + key
}
