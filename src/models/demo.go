/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file models user.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"demo/demo"
	"utils"
	"fmt"
)

var engine *xorm.Engine

const (
	APP_CONFIG = "/conf/db.conf"
)


type Demo struct {
	Demo_id   int64  `xorm:"demo_id pk autoincr"`
	Demo_name string `xorm:"demo_name varchar(100) not null"`

}

/**
 * 初始化方法
 **/
func init() {

	//初始化数据库连接
	engine, _ = utils.GetDB()
}

/**
*获取用户信息
 */

func GetDemoInfo(demoId int) ([]*demo.DemoInfo, error) {

	newDemo := make([]Demo, 0)
	errmu := engine.Sql("select * from demo where demo_id = ? ", demoId).Find(&newDemo)
	//errmu := engine.Table("user").In("user_id", 200, 201, 202).Find(&newUser)

	if errmu != nil {
		utils.LogErr("failed to get data info, input_param:%s", errmu)
	}
	
	fmt.Println(newDemo)

	demoList := make([]*demo.DemoInfo, len(newDemo))

	for k, v := range newDemo {
		demoInfo := new(demo.DemoInfo)
		demoInfo.DemoID = v.Demo_id
		demoInfo.DemoName = v.Demo_name
		demoList[k] = demoInfo
		demoInfo = nil
	}

	return demoList, errmu
}

/**
* 新创建用户
 */
func CreateNewDemo(paramMap map[string]string) (int64, int64, error) {
	
	err := engine.Sync2(new(Demo))
	if err != nil{
		
	}

	demo := new(Demo)
	demo.Demo_name = paramMap["demo_name"]
	intAffected, err := engine.Insert(demo)

	return demo.Demo_id, intAffected, err

}
