// Paceage user
/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-06 10:31:18
 * @LastEditTime: 2019-11-06 15:43:38
 * @LastEditors: Dawn
 */

package user

import (
	"fmt"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/wangcong0918/sunrise"
	"github.com/wangcong0918/sunrise/log"
	"github.com/wangcong0918/sunrise/utils/jwt"
	"os"
	"time"
	"userAuthority/api/maps"
	"userAuthority/api/middleware"
	"userAuthority/api/models"
	"userAuthority/api/thirdUtils"
)

// 用户登录
func WeChatAppletLogin(t *sunrise.Context) {
	value, exist := t.Get("WeChatAppletLogin") //获取绑定数据
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	var userInfo models.S_user
	var openID = ""
	var tokenString = ""
	var appId, secret string
	login := value.(maps.WeChatAppletLoginRequest)
	if login.WeChatAppletType == "1" { //驾驶企业端
		appId = os.Getenv("QJAPPID")
		secret = os.Getenv("QJSECRET")
	} else if login.WeChatAppletType == "2" { //销售端
		appId = os.Getenv("XSAPPID")
		secret = os.Getenv("XSSECRET")
	}
	session, errs := wechat.Code2Session(appId, secret, login.Code)
	if errs != nil || session.Errcode > 0 {
		log.Logger.Info("Code2Session------------->", errs)
		resp.Code = 94
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	openID = session.Openid
	phone := new(wechat.UserPhone)
	err := wechat.DecryptOpenDataToStruct(login.EncryptedData, login.Iv, session.SessionKey, phone)
	if err != nil {
		log.Logger.Info("DecryptOpenDataToStruct------------->", err)
		resp.Code = 99
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	//判断用户是否存在
	userInfo = models.GetUserInfoByLoginName(phone.PhoneNumber)
	if len(userInfo.UserID) > 1 {
		jwtInfo := new(jwt.User)
		jwtInfo.UserID = userInfo.UserID
		jwtInfo.OpenID = openID
		token, jwtErr := jwt.JwtGenerateToken(jwtInfo, 720*time.Hour)
		if jwtErr != nil {
			resp.Code = 2
			resp.Msg = maps.Msg[resp.Code]
			middleware.ResponseResult(t, resp)
			return
		}
		tokenString = token
		thirdUtils.RedisDelAndSet(userInfo.UserID, tokenString) //添加新的token 有效期一个月
	} else {
		resp.Code = 9
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	var authorization maps.Authorization
	authorization.Authorization = tokenString
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = authorization
	middleware.ResponseResult(t, resp)
	return
}

// 修改密码
func UpdatePwd(t *sunrise.Context) {
	value, exist := t.Get("UpdatePwd")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	updatePwd := value.(maps.UpdatePwd)
	msg, code := models.UpdatePwd(updatePwd)
	resp.Code = code
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = msg
	middleware.ResponseResult(t, resp)
	return
}

// 新增用户
func InsertUser(t *sunrise.Context) {
	value, exist := t.Get("InsertUser")
	var resp maps.ResponseInfo
	var loginInfo models.S_user
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	insertUser := value.(maps.InsertUser)
	// 判断下 isSysOrApp 0/1
	//var isApp int64 = 0
	//if insertUser.IsSysOrApp == &isApp {
	//	loginInfo, _ = models.GetUserInfoByLoginName(insertUser.LoginName)
	//} else {
	//	loginInfo, _ = models.GetUserInfoByTel(insertUser.LoginName)
	//}

	if len(loginInfo.UserID) > 0 {
		resp.Code = 6
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	msg, userID := models.InsertUser(insertUser)
	if len(userID) > 0 {
		for _, value := range insertUser.UserRoleInfoList {
			value.UserID = userID
			models.InsertUserRole(value)
		}
		for _, value := range insertUser.UserAreaInfoList {
			value.UserID = userID
			models.InsertUserArea(value)
		}
	}
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = msg
	middleware.ResponseResult(t, resp)
	return
}

// 修改用户
func UpdateUser(t *sunrise.Context) {
	value, exist := t.Get("UpdateUser")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	updateUser := value.(maps.UpdateUser)
	msg := models.UpdateUser(updateUser)
	msgdel := models.DelUserAreaByUserID(updateUser.UserID)
	msgdels := models.DelUserRoleByUserID(updateUser.UserID)
	fmt.Println(msgdels, "DelUserRoleByUserID返回结果！")
	fmt.Println(msgdel, "DelUserAreaByUserID返回结果！")

	for _, value := range updateUser.UserRoleInfoList {
		value.UserID = updateUser.UserID
		models.InsertUserRole(value)
	}
	for _, value := range updateUser.UserAreaInfoList {
		value.UserID = updateUser.UserID
		msg := models.InsertUserArea(value)
		fmt.Println(value, msg, "InsertUserArea返回结果！")
	}
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = msg
	middleware.ResponseResult(t, resp)
	return
}

// 删除用户
func UpdateUserIsDelete(t *sunrise.Context) {
	value, exist := t.Get("UpdateUserIsDelete")
	var resp maps.ResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	delUser := value.(maps.UpdateUserIsDelete)
	models.UpdateUserIsDelete(delUser.UserID)
	models.DelUserAreaByUserID(delUser.UserID)
	models.DelUserRoleByUserID(delUser.UserID)
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = "删除成功"
	middleware.ResponseResult(t, resp)
	return
}

// 动态获取所有用户
func GetUserListInfo(t *sunrise.Context) {
	value, exist := t.Get("GetUserListInfo")
	var resp maps.ResponseInfo
	var re maps.GetUserListResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	getUserListInfo := value.(maps.GetUserListInfo)
	count := models.GetUserListInfoCount(getUserListInfo)
	if count > 0 {
		userList := models.GetUserListInfo(getUserListInfo)
		re.UserList = userList
	}

	re.Count = count
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = re
	middleware.ResponseResult(t, resp)
	return
}

//用户退出
func OutLogin(t *sunrise.Context) {
	v, existKey := t.Get("contextJwtUserInfo")
	var resp maps.ResponseInfo
	if !existKey {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	userInfo := v.(jwt.User)
	thirdUtils.RedisDel(userInfo.UserID)
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = "退出成功！"
	middleware.ResponseResult(t, resp)
	return
}

// 获取系统日志
func GetSystemLogList(t *sunrise.Context) {
	value, exist := t.Get("GetSystemLogList")
	var resp maps.ResponseInfo
	var re maps.GetSystemLogListResponseInfo
	if !exist {
		resp.Code = 1
		resp.Msg = maps.Msg[resp.Code]
		middleware.ResponseResult(t, resp)
		return
	}
	getSystemLogList := value.(maps.GetSystemLogListInfo)
	sysList, count, _ := models.GetSystemLogList(getSystemLogList)
	re.SysList = sysList
	re.Count = count
	resp.Code = 0
	resp.Msg = maps.Msg[resp.Code]
	resp.Result = re
	middleware.ResponseResult(t, resp)
	return
}
