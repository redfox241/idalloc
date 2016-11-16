package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	//"strconv"
	"time"
	"idalloc/idalloc"
)

const (
	HOST = "127.0.0.1"
	PORT = "9090"
)

func main() {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := idalloc.NewIdallocClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	//	model := user.UserInfo{
	//		100,
	//		"诸葛亮",
	//		"孔明",
	//		"鞠躬尽瘁，死而后已",
	//	}

	intIndex := 10000
	intStep := 0
	
	for  intStep < intIndex {
		
		
		model := make(map[string]string)
		model["id"] = "0"
		model["type_name"] = "msg_id"
	
	
		intNewDemoId , _ := client.GenId(model)
	
		fmt.Println("new_msg_id : ", intNewDemoId)
		
		intStep ++
		
	}

	endTime := currentTimeMillis()
	fmt.Printf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
