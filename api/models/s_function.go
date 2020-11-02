/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-09-02 15:09:00
 * @LastEditors: Dawn
 */
package models

import (
	"fmt"
	"time"
)

type S_function struct {
	FunctionID         *string   `json:"functionID"  xorm:"notnull 'functionID' default('')"`         // guid
	SystemCode         *string   `json:"systemCode"  xorm:"'systemCode' default()"`                   // 系统编码
	ParentFunctionCode *string   `json:"parentFunctionCode"  xorm:"'parentFunctionCode' default('')"` // 功能编码父节点
	FunctionCode       *string   `json:"functionCode"  xorm:"notnull 'functionCode' default()"`       // 功能编码
	FunctionName       *string   `json:"functionName"  xorm:"'functionName' default('')"`             // 功能名称
	FunctionType       *int64    `json:"functionType"  xorm:"'functionType' default(0)"`              // 功能类型(0菜单 1按钮)
	FunctionMethod     *string   `json:"functionMethod"  xorm:"'functionMethod' default()"`           // 请求方法:get、post、delete、put
	FunctionIcon       *string   `json:"functionIcon"  xorm:"'functionIcon' default()"`               // 菜单图标
	Description        *string   `json:"description"  xorm:"'description' default()"`                 // 功能描述
	Index              *int64    `json:"index"  xorm:"'index' default(0)"`                            // 排序号
	IsValid            *int64    `json:"isValid"  xorm:"'isValid' default(0)"`                        // 是否有效 0-有效 1-无效 0-有效  1-无效
	Path               *string   `json:"path"  xorm:"'path' default()"`                               // 权限路径
	InsertTime         time.Time `json:"insertTime"  xorm:"'insertTime' default(CURRENT_TIMESTAMP)"`  // 记录插入时间
	UpdateTime         time.Time `json:"updateTime"  xorm:"'updateTime' default(CURRENT_TIMESTAMP)"`  // 记录更新时间
}

func (function S_function) TableName() string {
	return "s_function"
}

// 根据userID账号查看是否有权限
func GetFunctionByUserID(userID string, requestURI string) bool {
	sql := ` SELECT * from s_function WHERE functionID in(
		 SELECT functionID from  s_role_function WHERE roleID =(SELECT roleID from s_user_role where s_user_role.userID=?)
		 )  and s_function.path=? `
	auth, err := Engine.QueryString(sql, userID, requestURI)
	if len(auth) <= 0 || err != nil {
		fmt.Println(err, "GetFunctionByUserID失败原因")
		return false
	}
	if len(auth) > 0 {
		return true
	}
	return false
}
