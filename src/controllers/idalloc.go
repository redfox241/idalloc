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
	
	//newIdId, err := models.CreateNewIdalloc(idallocInfo)
	newIdId, err := models.AllocId(idallocInfo)
	
	//newIdId,intAffects,err := models.CreateNewIdFormFile(idallocInfo)
	if newIdId > 0 {
		return newIdId, err
	} else {
		return newIdId, err
	}

}
