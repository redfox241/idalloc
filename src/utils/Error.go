/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file utils Conf.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package utils

func CheckErr(err error){
	
	LogErr("system hed error:",err)
	if err != nil{
		panic(err)
	}
}
