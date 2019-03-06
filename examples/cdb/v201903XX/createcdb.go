package main

import (
	"encoding/json"
	"fmt"
	"os"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

func main() {
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 5
	//cpf.HttpProfile.Endpoint = "cvm.ap-guangzhou.tencentcloudapi.com"
	cpf.SignMethod = "HmacSHA1"

	client, _ := cdb.NewClient(credential, "na-siliconvalley", cpf)
	request := cdb.NewCreateDBInstanceHourRequest()
	request.GoodsNum = common.Int64Ptr(1)
	request.Memory = common.Int64Ptr(1000)
	request.Volume = common.Int64Ptr(25)
	request.Zone = common.StringPtr("na-siliconvalley-2")
	response, err := client.CreateDBInstanceHour(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// unexpected errors
	// if err != nil {
	// 	panic(err)
	// }
	b, _ := json.Marshal(response.Response)
	fmt.Printf("%s", b)
}
