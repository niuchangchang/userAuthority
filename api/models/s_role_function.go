/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-09-29 10:21:58
 * @LastEditors: Dawn
 */
package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"userAuthority/api/maps"
	"userAuthority/api/thirdUtils"
)

type S_role_function struct {
	RfID       string    `json:"rfID"  xorm:"notnull 'rfID' default('')"`             // 权限关系标识
	RoleID     string    `json:"roleID"  xorm:"notnull 'roleID' default('')"`         // 角色标识
	FunctionID string    `json:"functionID"  xorm:"notnull 'functionID' default('')"` // 功能标示
	InsertTime time.Time `json:"insertTime" xorm:"created 'insertTime'"`              // 记录插入时间
	UpdateTime time.Time `json:"updateTime" xorm:"updated 'updateTime'"`              // 记录更新时间
}

func (role_function S_role_function) TableName() string {
	return "s_role_function"
}

// 新增角色权限
func InsertRoleFunction(RoleFunctionInfo maps.RoleFunctionInfo) (msg string) {
	RoleFunctionInsertInfo := new(S_role_function)
	b, _ := json.Marshal(RoleFunctionInfo)
	json.Unmarshal(b, &RoleFunctionInsertInfo)
	uuid := thirdUtils.UUID()
	RoleFunctionInsertInfo.RfID = uuid
	affected, err := Engine.Insert(RoleFunctionInsertInfo)
	if err != nil {
		fmt.Println(err, "InsertRoleFunction失败原因")
		return "新增失败！"
	} else {
		if affected > 0 {
			return "新增成功！"
		} else {
			return "新增失败！"
		}
	}
}
func DelRoleFunctionByRoleID(roleID string) (msg string) {
	if exist, err := Engine.Exec("DELETE FROM s_role_function WHERE roleID = ?", roleID); err != nil {
		fmt.Println(err, "DelRoleFunctionByUserID失败原因")
		return "删除失败！"
	} else {
		info, _ := exist.RowsAffected()
		if info > 0 {
			return "删除成功！"
		} else {
			return "删除失败！"
		}
	}
}
func GetRoleFunctionByRoleID(roleID, systemCode string) (functionList []maps.FunctionInfo, err error) {
	var sqlStr = ""
	var num int64 = 0
	var roleFunction []map[string]string
	if len(roleID) <= 0 {
		sqlStr += `SELECT * from s_function where s_function.parentFunctionCode=0 ` //ORDER BY s_function.index ASC
		roleFunction, err = Engine.QueryString(sqlStr)
		num = 1
	} else {
		sqlStr += ` SELECT s_function.* from s_role_function LEFT JOIN s_function on s_function.functionID=s_role_function.functionID WHERE s_role_function.roleID = ? and s_function.parentFunctionCode=0  ` //ORDER BY s_function.index ASC
		num = 2
		roleFunction, err = Engine.QueryString(sqlStr, roleID)
	}
	b, _ := json.Marshal(roleFunction)
	json.Unmarshal(b, &functionList)

	for k, value := range functionList {
		functionList[k].Children, err = Recursive(value.FunctionCode, roleID, &(functionList[k].Children), num)
		fmt.Println(functionList[k].Children)
	}

	if err != nil {
		fmt.Println(err, "GetRoleFunctionByRoleID失败原因")
		return functionList, err
	}
	return functionList, nil
}

func Recursive(functionCode, roleID string, children *[]maps.FunctionInfo, num int64) (funtionList []maps.FunctionInfo, err error) {
	var sqlStr = ""
	var roleFunction []map[string]string
	if num == 1 {
		sqlStr += `SELECT * from s_function WHERE parentFunctionCode=?` // ORDER BY s_function.index ASC
		roleFunction, err = Engine.QueryString(sqlStr, functionCode)
	} else {
		sqlStr += `SELECT s_function.* from s_role_function LEFT JOIN s_function on s_function.functionID=s_role_function.functionID WHERE s_role_function.roleID = ? and s_function.parentFunctionCode=? ` //ORDER BY s_function.index ASC
		roleFunction, err = Engine.QueryString(sqlStr, roleID, functionCode)
	}

	for key, value := range roleFunction {
		FunctionID := value["functionID"]
		FunctionCode := value["functionCode"]
		ParentFunctionCode := value["parentFunctionCode"]
		FunctionName := value["functionName"]
		FunctionType, _ := strconv.ParseInt(value["functionType"], 10, 64)
		FunctionMethod := value["functionMethod"]
		FunctionIcon := value["functionIcon"]
		Description := value["description"]
		IsValid, _ := strconv.ParseInt(value["isValid"], 10, 64)
		Path := value["path"]
		var tmpVal maps.FunctionInfo
		tmpVal.FunctionID = FunctionID
		tmpVal.ParentFunctionCode = ParentFunctionCode
		tmpVal.FunctionCode = FunctionCode
		tmpVal.FunctionName = FunctionName
		tmpVal.FunctionType = FunctionType
		tmpVal.FunctionMethod = FunctionMethod
		tmpVal.FunctionIcon = FunctionIcon
		tmpVal.Description = Description
		tmpVal.IsValid = IsValid
		tmpVal.Path = Path
		*children = append(*children, tmpVal)
		realChildren := *children
		Recursive(FunctionCode, roleID, &(realChildren[key].Children), num)

	}
	return *children, nil
}
