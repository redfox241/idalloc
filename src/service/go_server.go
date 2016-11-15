/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file main
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package main

import (
	"controllers"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"os"
	"demo/demo" //注意导入Thrift生成的接口包
	"utils"
)

const (
	APP_CONFIG = "/conf/app.conf"
)

func main() {

	//get conf
	utils.SetConfInfo(APP_CONFIG)
	ip_addr := utils.GetValuesByKeys("server_setting","ip_addr").(string)
	port := utils.GetValuesByKeys("server_setting","port").(string)

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(ip_addr+":"+port)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := controllers.GetDemoThrift()
	processor := demo.NewDemoProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	utils.LogDebug("Starting the simple server on :", ip_addr+":"+port)
	server.Serve()

}
