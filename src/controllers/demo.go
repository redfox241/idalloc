/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file demo.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package controllers

import (
	"models"
	"strconv"
	"demo/demo"
	"utils"
)

type demoThrift struct {
}

func GetDemoThrift() *demoThrift {
	return &demoThrift{}
}

func (this *demoThrift) GetDemoInfo(paramMap map[string]string) ([]*demo.DemoInfo, error) {

	var err error
	
	intDemoId, _ := strconv.Atoi(paramMap["demo_id"])
	newDemo := make([]*demo.DemoInfo, 0)
	newDemo, err = models.GetDemoInfo(intDemoId)

	utils.LogErr("-->from client Call:", paramMap)
	utils.LogNotice("finish to get demo info.")

	return newDemo, err
}


func (this *demoThrift) ProcessDemo(paramMap map[string]string) (int64, error) {

	demoinfo := make(map[string]string)
	demoinfo["demo_id"] = "0"
	demoinfo["demo_name"] = paramMap["demo_name"]
	
	newDemoId, intAffects, err := models.CreateNewDemo(demoinfo)
	if intAffects > 0 {
		return newDemoId, err
	} else {
		return newDemoId, err
	}

}
