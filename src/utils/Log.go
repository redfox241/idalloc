/***************************************************************************
 *
 * Copyright (c) 2016 primedu.com, Inc. All Rights Reserved
 *
 **************************************************************************/

/**
 * @file utils log.go
 * @author bugushe@gmail.com
 * @date 2016-10-15 13:50:37
 * @brief
 *
 **/

package utils

import (
	"log"
	"os"
)

var strLogFileName, strPathName string

func init() {

	strLogFileName = createFolderExist("") + "/server.log"
	strPathName = strLogFileName

}

func LogFatal(content ...interface{}) {

	appendToFile(strLogFileName, "FATAl", content)
}

func LogErr(content ...interface{}) {

	appendToFile(strLogFileName, "ERROR", content)
}

func LogInfo(content ...interface{}) {
	appendToFile(strLogFileName, "INFO", content)
}

func LogDebug(content ...interface{}) {

	appendToFile(strLogFileName, "DEBUG", content)
}

func LogNotice(content ...interface{}) {
	appendToFile(strLogFileName, "NOTICE", content)
}

/*
 fileName:文件名字(带全路径)
 content: 写入的内容
*/
func appendToFile(fileName string, logType string, content ...interface{}) error {

	logfile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln("open file error !")
	}

	logger := log.New(logfile, "\r\n", log.LstdFlags|log.Llongfile)
	logger.SetPrefix("[" + logType + "]")

	logger.Println(content)

	defer logfile.Close()

	return err
}

/*
* 创建文件夹
 */
func createFolderExist(strFolder string) string {

	//设定文件夹默认值 logs
	if len(strFolder) == 0 {
		strFolder = "logs"
	}

	var strNewPathName = getParentDirectory(getCurrentDirectory()) + "/" + strFolder

	if _, err := os.Stat(strNewPathName); os.IsNotExist(err) {
		newerr := os.MkdirAll(strNewPathName, 0777)
		if newerr != nil {
			os.MkdirAll(strNewPathName, 0777)
		}
	}
	return strNewPathName
}
