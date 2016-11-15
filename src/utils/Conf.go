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

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

var confPath string
var confContent []byte

func SetConfInfo(filePath string) {
	
	confPath = filePath
	
}

//解析文件，取出所有参数
func GetAppConfig(filePath string) map[interface{}]interface{} {
	var strAppPath, strFileName string
	var err error
	
	strAppPath = getParentDirectory(getCurrentDirectory())
	strFileName = strAppPath + "/" + filePath
	
	confContent, err = ioutil.ReadFile(strFileName)
	
	//将解析出的参数转为map的形式
	mapAppConf := make(map[interface{}]interface{})
	if err != nil {
		LogErr("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(confContent), &mapAppConf)
	return mapAppConf
}

//根据需求取出对应值
func GetElement(key string, themap map[interface{}]interface{}) string {
	
	if value, ok := themap[key]; ok {
		return value.(string)
	}
	LogErr("Can't find the *.yaml")
	return ""
}

func GetValByKey(key string, themap map[interface{}]interface{}) interface{} {
	
	if _, ok := themap[key]; ok {
		
		switch reflect.TypeOf(themap[key]).Kind() {
		case reflect.Array:
			return themap[key].(map[interface{}]interface{})
		//return "array"
		case reflect.Map:
			return themap[key].(map[interface{}]interface{})
		//return "map"
		case reflect.Bool:
			return themap[key]
		case reflect.Int:
			return themap[key]
		case reflect.String:
			return themap[key]
		case reflect.Float64:
			return themap[key]
		default:
			fmt.Println("not match type")
		}
		
		return nil
	}
	
	LogErr("Can't find the key ,key:", key)
	return nil
}

func GetValuesByKeys(keys ...string) interface{} {
	
	var res bool
	appConfContent := make(map[interface{}]interface{})
	appConfContent = GetAppConfig(confPath)
	
	intKeys := len(keys)
	if intKeys <= 1 {
		
		if _, res = appConfContent[keys[0]]; ! res {
			fmt.Println("the key is not exist,key:%s", keys[0])
		}
		
		return appConfContent[keys[0]]
	} else {
		
		confContent := GetValByKey(keys[0], appConfContent)
		for intIndex := 1; intIndex < len(keys); intIndex ++ {
			switch reflect.TypeOf(confContent).Kind() {
			case reflect.Map:
				confContent = GetValByKey(keys[intIndex], confContent.(map[interface{}]interface{}))
			case reflect.Array:
				fmt.Println(reflect.TypeOf(confContent))
				confContent = GetValByKey(keys[intIndex], confContent.(map[interface{}]interface{}))
			default:
				confContent = GetValByKey(keys[intIndex], confContent.(map[interface{}]interface{}))
			}
		}
		
		return confContent
	}
	
	LogErr("Can't find the key ,key:[%s]", keys)
	return nil
}