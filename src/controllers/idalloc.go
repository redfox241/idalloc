/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file idalloc.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package controllers

import (
	"models"
	"utils"
	"errors"
)

const (
	APP_CONFIG = "/conf/app.conf"
)

type idallocThrift struct {
}

func GetIdallocThrift() *idallocThrift {
	return &idallocThrift{}
}

func (this *idallocThrift) GenId(paramMap map[string]string) (int64, error) {
	
	idallocInfo := make(map[string]string)
	idallocInfo["idalloc_id"] = "0"
	idallocInfo["type_name"] = paramMap["type_name"]
	
	//read conf for get save_mode
	//get conf
	utils.SetConfInfo(APP_CONFIG)
	saveMode := utils.GetValuesByKeys("idalloc_info", "save_mode").(int)
	
	utils.LogInfo("save mode :", saveMode)
	
	var newIdId int64 = 0
	var err error = errors.New("EXIT")
	
	//支持多模式
	switch saveMode {
		// defalut mysql
		case 0 :
			newIdId, err = models.AllocIdByMysql(idallocInfo)
		// redis
		case 1 :
			newIdId, err = models.AllocIdByRedis(idallocInfo)
		//file
		case 2 :
		
		//file
		default :
		
	}
	
	if newIdId > 0 {
		return newIdId, err
	} else {
		return newIdId, err
	}
	
}
