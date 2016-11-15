/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file utils db.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/
package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/go-xorm/xorm"
)

var dbconnect *Engine

const (
	DB_CONFIG = "/conf/db.conf"
)

func GetDB() (*Engine, error) {

	if dbconnect != nil {

		return dbconnect, nil
	} else {
		LogDebug("db connect init ...")
	}

	//get conf
	SetConfInfo(DB_CONFIG)
	db_driver := GetValuesByKeys("db_driver").(string)
	db_name := GetValuesByKeys("db_name").(string)
	user := GetValuesByKeys("user").(string)
	passwd := GetValuesByKeys("passwd").(string)
	ipaddr := GetValuesByKeys("ip_addr").(string)
	port := GetValuesByKeys("port").(string)
	db_charset := GetValuesByKeys("charset").(string)

	var err error
	strConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, passwd, ipaddr, port, db_name, db_charset)
	//fmt.Println(strConnect)
	dbconnect, err = NewEngine(db_driver, strConnect)
	dbconnect.ShowSQL(true)

	if err != nil {
		LogErr(fmt.Printf("failed to connect mysql.error:%s", err))
	}
	return dbconnect, err
}
