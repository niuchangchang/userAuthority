/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-21 15:16:18
 * @LastEditTime: 2019-09-03 14:55:38
 * @LastEditors: Dawn
 */
package user

import (
	"github.com/wangcong0918/sunrise"
	"github.com/wangcong0918/sunrise/utils/jwt"
	"os"
	"userAuthority/api/maps"
	"userAuthority/api/middleware"
	"userAuthority/api/models"
)

// 获取所有角色
func GetRoleList(t *sunrise.Context) {
	var resp maps.ResponseInfo
	roleList, _ := models.GetRoleList()
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = roleList
	middleware.ResponseResult(t, resp)
	return
}

// 新增角色
func InsertRole(t *sunrise.Context) {
	value, exist := t.Get("InsertRole")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	insertRole := value.(models.S_role)
	roleInfo, _ := models.GetRoleInfoByRoleName(insertRole.RoleName)
	if roleInfo {
		resp.Code = 6
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	msg := models.InsertRole(insertRole)
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = msg
	middleware.ResponseResult(t, resp)
	return
}

//修改角色
func UpdateRole(t *sunrise.Context) {
	value, exist := t.Get("UpdateRole")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	updateRole := value.(models.S_role)
	msg := models.UpdateRole(updateRole)
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = msg
	middleware.ResponseResult(t, resp)
	return
}

//删除角色
func DeleteRoleByRoleID(t *sunrise.Context) {
	value, exist := t.Get("DeleteRoleByRoleID")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	delRole := value.(maps.DeleteRoleByRoleID)

	// 默认的管理员账号不可删
	defaultMessageCount := os.Getenv("DEFAULT_MESSAGE_COUNT")
	if delRole.RoleID == defaultMessageCount {
		resp.Code = 1
		resp.Msg = "管理员账号不可删除"
		middleware.ResponseResult(t, resp)
		return
	}

	msg := models.DeleteRoleByRoleID(delRole.RoleID)
	models.DelRoleFunctionByRoleID(delRole.RoleID)
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = msg
	middleware.ResponseResult(t, resp)
	return
}

// 获取所有角色权限
func GetRoleFunctionByRoleID(t *sunrise.Context) {
	v, existKey := t.Get("contextJwtUserInfo")
	value, exist := t.Get("GetRoleFunctionByRoleID") //获取绑定数据
	var resp maps.ResponseInfo
	if !exist || !existKey {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	userInfo := v.(jwt.User)
	roleRequest := value.(maps.GetRoleFunctionByRoleIDRequest)
	roleInfo := models.GetRoleIDByUserID(userInfo.UserID)
	roleFunctionList, _ := models.GetRoleFunctionByRoleID(roleInfo.RoleID, roleRequest.Platform)
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = roleFunctionList
	middleware.ResponseResult(t, resp)
	return
}

// 新增角色权限
func InsertRoleFunction(t *sunrise.Context) {
	value, exist := t.Get("InsertRoleFunction")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	insertRoleFunction := value.(maps.InsertRoleFunctionInfo)
	models.DelRoleFunctionByRoleID(insertRoleFunction.RoleFunctionInfoList[0].RoleID)
	for _, value := range insertRoleFunction.RoleFunctionInfoList {
		models.InsertRoleFunction(value)
	}
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = "成功！"
	middleware.ResponseResult(t, resp)
	return
}
