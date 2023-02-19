package obss

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
)

func putFile(key string, source string) {
	input := &obs.PutFileInput{}
	input.Bucket = bucketName
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
}

func genGetURL(key string) string {
	return "https://" + bucketName + "." + endPoint + "/" + key
}
