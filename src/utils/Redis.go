/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file utils redis.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/
package utils

import (
	"github.com/garyburd/redigo/redis"
)

var redisconnect redis.Conn = nil

const (
	REDIS_CONFIG = "/conf/db.conf"
)

func init () {
	
}

/**
 * 连接redis
 */
func GetRedisConnect() ( redis.Conn, error) {

	if redisconnect != nil {
		return redisconnect,nil
	}else{
		LogDebug("success to init redis connect")
	}
	
	//get conf
	SetConfInfo(REDIS_CONFIG)
	protocol := GetValuesByKeys("redis_setting","redis_protocol").(string)
	host := GetValuesByKeys("redis_setting","redis_host").(string)
	port := GetValuesByKeys("redis_setting","redis_port").(string)
	//pwd := GetValuesByKeys("redis_setting","redis_pwd").(string)
	//prefix := GetValuesByKeys("redis_setting","redis_key_prefix").(string)
	
	
	redisconnect, err := redis.Dial( protocol , host + ":" + port )
	
	if err != nil {
		LogErr("failed to connect redis,err:",err)
		CheckErr(err)
	}
	//defer redisconnect.Close()

	return redisconnect,err

}
