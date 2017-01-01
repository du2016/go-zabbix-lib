package main

import (
	"fmt"
	zabbix "github.com/du2016/go-zabbix-lib"
	"log"
)

func main() {
	//创建API实例
	api := zabbix.NewAPI("http://127.0.0.1/api_jsonrpc.php")
	//登录
	api.Login("admin", "zabbix")
	//退出
	defer api.Logout()
	//更改对影类别action状态
	api.ActionUpdatestatusByStatus("0", "0", "0")
	//根据hostname获取host实例
	host, err := api.HostGetByHost("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	//获取host对应的ID
	for _, v := range host.Groups {
		fmt.Printf("%#v\n", v)
	}
}
