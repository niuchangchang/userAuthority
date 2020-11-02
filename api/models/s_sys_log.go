/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-10-22 10:18:27
 * @LastEditors: Dawn
 */
package models

import (
	"fmt"
	"time"
	"userAuthority/api/maps"
	"userAuthority/api/thirdUtils"
)

type S_sys_log struct {
	LogID        string    `json:"logID"  xorm:"notnull 'logID' default('')"`       // 日志标识
	LocationCode string    `json:"locationCode"  xorm:"'locationCode' default('')"` // 地址编码
	ReqUrl       string    `json:"reqUrl"  xorm:"notnull 'reqUrl' default()"`       // http请求地址
	ReqMethod    string    `json:"reqMethod"  xorm:"notnull 'reqMethod' default()"` // http请求方法
	UserName     string    `json:"userName"  xorm:"'userName' default()"`           // 登录用户
	UserIp       string    `json:"userIp"  xorm:"'userIp' default()"`               // 请求ip
	IpRegion     string    `json:"ipRegion"  xorm:"'ipRegion' default('')"`         // ip地址归属地
	HandleMethod string    `json:"handleMethod"  xorm:"'handleMethod' default()"`   // 处理方法
	ReqArgs      string    `json:"reqArgs"  xorm:"notnull 'reqArgs' default('')"`   // 请求参数
	BrowserInfo  string    `json:"browserInfo"  xorm:"'browserInfo' default('')"`   // 浏览器信息
	Result       string    `json:"result"  xorm:"'result' default('')"`             // 浏览器信息
	ReqTime      time.Time `json:"reqTime"  xorm:"'created' 'reqTime' "`            // 处理时间
}

func (sys_log S_sys_log) TableName() string {
	return "s_sys_log"
}

// 新增日志
func InsertSysLog(insertSysLog S_sys_log) {
	uuid := thirdUtils.UUID()
	insertSysLog.LogID = uuid
	affected, err := Engine.Insert(insertSysLog)
	if err != nil {
		fmt.Println(err, "InsertSysLog失败原因")
	} else {
		if affected > 0 {
			fmt.Println("InsertSysLog成功")
		} else {
			fmt.Println(err, "InsertSysLog失败原因")
		}
	}
}

// 查询所有日志
func GetSystemLogList(info maps.GetSystemLogListInfo) (logList []map[string]string, count int64, err error) {
	startIndex := (*info.PageNum - 1) * *info.PageSize
	pageSize := *info.PageSize
	sql := "select * FROM s_sys_log order by reqTime desc limit ?,? "
	logLists, errs := Engine.QueryString(sql, startIndex, pageSize)
	sqlcount := "select * FROM s_sys_log"
	counts, err := Engine.QueryString(sqlcount)
	countss := len(counts)
	if err != nil || errs != nil {
		fmt.Println(err, "GetSystemLogList失败原因")
		return logLists, 0, err
	}
	return logLists, int64(countss), nil
}

//func InsertDataLog(result string, t *sunrise.Context) {
//	//存入登录日志
//	token := t.GetHeader("Authorization")
//	client, clienterr := thirdUtils.NewClient(10)
//	fmt.Println(clienterr, "Authorization打开redis错误原因")
//	val, err := client.Get(token).Result()
//	var userInfo S_user
//	Engine.Where("userID = ? ", val).Get(&userInfo)
//	fmt.Println(val, "val------------>", err, "err------------>")
//	ip := t.ClientIP()
//	rebody := t.Request
//	reqUrl := rebody.Host + rebody.RequestURI
//	insertSysLog := new(S_sys_log)
//	insertSysLog.UserIp = ip
//	insertSysLog.ReqUrl = reqUrl
//	insertSysLog.ReqMethod = rebody.Method
//	insertSysLog.BrowserInfo = t.GetHeader("User-Agent")
//	insertSysLog.HandleMethod = rebody.RequestURI
//	insertSysLog.ReqTime = time.Now()
//	insertSysLog.UserName = *userInfo.LoginName
//	insertSysLog.Result = result
//	//插入登录日志
//	affected, err := Engine.Insert(insertSysLog)
//	if err != nil {
//		fmt.Println(err, "InsertDataLog失败原因")
//	} else {
//		if affected > 0 {
//			fmt.Println("InsertDataLog成功")
//		} else {
//			fmt.Println(err, "InsertDataLog失败原因")
//		}
//	}
//}
