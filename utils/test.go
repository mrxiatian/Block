package utils

import (
	"bitcoin/entity"
	"fmt"
	"strings"
)


//获取区块链信息
func GetTest() {
	paramsSlice := []interface{}{"xiatian"}
	//RPC通信标椎格JSON式数据
	rpcNormJson := PrepareJSON("getaddressesbylabel", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := DoPost(RPCURL, RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	label := entity.Label{}
	//fmt.Printf("%T %v", rpcResult.Data.Result, rpcResult.Data.Result)
	fmt.Printf("")
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {

		for _, v1 := range res {
			res1, ok := v1.(map[string]interface{})
			if ok {
				for _, v2 := range res1 {
					str, ok := v2.(string)
					if ok {
						add := entity.Address{}
						add.Purpose = str
						label.Address = append(label.Address, add)
					}
				}
			}

		}
		fmt.Println(label)
	}

	//return blockStats
}
