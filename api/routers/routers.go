// Package routers ...
/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-06 10:31:18
 * @LastEditTime: 2019-11-07 11:21:15
 * @LastEditors: Dawn
 */
package routers

import (
	// "os/user"
	//"userAuthority/api/controllers/area"
	//"userAuthority/api/controllers/device"
	"github.com/wangcong0918/sunrise"
	"userAuthority/api/controllers/user"
	"userAuthority/api/middleware"
)

// VERSION
var (
	VERSION = "/api/userAuthority/v1"
)

// User make routers 用户路由
func User(e *sunrise.Engine) *sunrise.RouterGroup {
	base := e.Group(VERSION + "/user")
	userService := base
	//用户权限验证
	userService.POST("/validateJwtMiddleware", middleware.ValidateJwtMiddleware)
	// 销售端小程序登录
	userService.POST("/weChatAppletLogin", middleware.CheckValidate, user.WeChatAppletLogin)

	userService.Use(middleware.ValidateJwtMiddleware)
	// 用户退出
	userService.POST("/outLogin", user.OutLogin)
	// 获取角色权限列表
	userService.POST("/getRoleFunctionByRoleID", middleware.CheckValidate, user.GetRoleFunctionByRoleID)

	//userService.Use(middleware.CheckValidate).Use(middleware.AuthToken)
	userService.POST("/updatePwd", user.UpdatePwd)                   // 修改密码
	userService.POST("/insertUser", user.InsertUser)                 // 新增用户
	userService.POST("/updateUser", user.UpdateUser)                 // 修改用户
	userService.POST("/updateUserIsDelete", user.UpdateUserIsDelete) // 删除用户
	userService.POST("/getUserListInfo", user.GetUserListInfo)       // 动态获取用户列表
	userService.POST("/getRoleList", user.GetRoleList)               // 获取角色列表
	userService.POST("/insertRole", user.InsertRole)                 // 新增角色
	userService.POST("/updateRole", user.UpdateRole)                 // 修改角色
	userService.POST("/deleteRoleByRoleID", user.DeleteRoleByRoleID) // 删除角色
	userService.POST("/insertRoleFunction", user.InsertRoleFunction) // 添加角色权限
	userService.POST("/getSystemLogList", user.GetSystemLogList)     // 获取日志列表

	return userService
}

//用户区域路由
func UserArea(e *sunrise.Engine) *sunrise.RouterGroup {
	base := e.Group(VERSION + "/userArea")
	userAreaService := base
	//userAreaService.Use(middleware.CheckValidate).Use(middleware.AuthToken)
	//userAreaService.POST("/insertUserArea", user.InsertUserArea) // 新增用户区域
	return userAreaService
}
