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
	"errors"
	"strings"
	"math/rand"
	"time"
)

var engine *xorm.Engine

const (
	APP_CONFIG = "/conf/app.conf"
)


type Idalloc struct {
	Id   int64  `xorm:"id pk autoincr"`
	Type_name string `xorm:"type_name varchar(100) not null"`
	Idalloc_id int64 `xorm:"idalloc_id bigint not null"`

}

/**
 * 初始化方法
 **/
func init() {

	//初始化数据库连接
	engine, _ = utils.GetDB()
}

/**
* 新创建id
 */
func CreateNewIdalloc(paramMap map[string]string) (int64, error) {
	
	//err := engine.Sync2(new(Idalloc))
	//if err != nil{
	//
	//}

	idInfo := new(Idalloc)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	idInfo.Idalloc_id = int64( r.Intn(100) )
	
	condInfo := new(Idalloc)
	condInfo.Type_name = strings.Trim( paramMap["type_name"],"")
	
	intAffected,err := engine.Update( idInfo,condInfo )
	
	fmt.Println(intAffected,err)

	return intAffected,err

}

func AllocId(paramMap map[string]string) (int64, error) {
	
	
	idallocInfo := make([]Idalloc, 0)
	var newIdallocId  int64 = 0
	
	session := engine.NewSession()
	defer session.Close()
	
	err := session.Begin()
	if err != nil{
		return 0,nil
	}
	
	_,err1 := session.Engine.Exec("update idalloc set idalloc_id = idalloc_id + 1 where type_name = ?",paramMap["type_name"])
	
	if err1 == nil{
		
		strSelectCmd := fmt.Sprintf("select * from idalloc where type_name = '%s' ",paramMap["type_name"] )
		errselect := session.Engine.Sql(strSelectCmd).Find(&idallocInfo)
		if errselect == nil{
			
		}
		
		for _,val := range idallocInfo {
			newIdallocId = (val.Idalloc_id)
		}
		
		session.Commit()
	}else{
		session.Rollback()
		
		return 0,err1
	}
	
	return  int64(newIdallocId),err1
	
}


func CreateNewIdFormFile( paramMap map[string] string ) (int64,int64,error){
	utils.SetConfInfo(APP_CONFIG)
	msg_id := utils.GetValuesByKeys("idalloc_info","msg_id").(int64)
	//user_id := utils.GetValuesByKeys("idalloc_info","user_id").(string)
	
	err := errors.New("EXIT")
	fmt.Println(msg_id)
	
	return msg_id,msg_id,err
}