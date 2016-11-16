/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file models idalloc.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"utils"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var engine *xorm.Engine
var rediscon redis.Conn

const (
	APP_CONFIG = "/conf/app.conf"
)

type Idalloc struct {
	Id         int64  `xorm:"id pk autoincr"`
	Type_name  string `xorm:"type_name varchar(100) not null"`
	Idalloc_id int64 `xorm:"idalloc_id bigint not null"`
}

/**
 * 初始化方法
 **/
func init() {
	
	//初始化数据库连接
	engine   , _  = utils.GetDB()
	rediscon , _  = utils.GetRedisConnect()
	
}


/*
 * 分配id，使用mysql存储
 *
 */

func AllocIdByMysql(paramMap map[string]string) (int64, error) {
	
	idallocInfo := make([]Idalloc, 0)
	var newIdallocId int64 = 0
	
	session := engine.NewSession()
	defer session.Close()
	
	err := session.Begin()
	if err != nil {
		utils.CheckErr(err)
	}
	
	strUpdate := fmt.Sprintf("insert into idalloc ( id,type_name,idalloc_id ) value (0,'%s','1')  " +
		"on duplicate key update idalloc_id = idalloc_id + 1 ", paramMap["type_name"])
	_, errUpdate := session.Engine.Exec(strUpdate)
	
	if errUpdate == nil {
		
		strSelectCmd := fmt.Sprintf("select * from idalloc where type_name = '%s' ", paramMap["type_name"])
		errget := session.Engine.Sql(strSelectCmd).Find(&idallocInfo)
		
		if errget != nil {
			utils.LogErr("failed to get id form idalloc")
			utils.CheckErr(errget)
		}
		
		for _, val := range idallocInfo {
			newIdallocId = (val.Idalloc_id)
		}
		
		session.Commit()
	} else {
		session.Rollback()
		utils.LogErr("failed to gen_id,error:", errUpdate)
		
		utils.CheckErr(errUpdate)
	}
	
	return int64(newIdallocId), errUpdate
	
}


/**
 * 分配id，by redis
 *
 */

func AllocIdByRedis(paramMap map[string]string) (int64, error) {
	
	strAllocKey := paramMap["type_name"]
	
	utils.SetConfInfo(APP_CONFIG)
	strKeyPrefix := utils.GetValuesByKeys("redis_key_prefix").(string)
	
	rediscon , _  = utils.GetRedisConnect()
	intVal, err := rediscon.Do("INCR", strKeyPrefix + strAllocKey )
	if err != nil {
		utils.LogErr("failed to incr redis,err:", err)
		utils.CheckErr(err)
	}
	
	return intVal.(int64),err
	
}