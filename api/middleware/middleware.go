/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-06 12:31:03
 * @LastEditTime: 2019-10-25 13:56:15
 * @LastEditors: Dawn
 */
package middleware

import (
	"fmt"
	//"JuFeng/leaseServer/thirdUtils"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/wangcong0918/sunrise"
	jwtInfo "github.com/wangcong0918/sunrise/utils/jwt"
	"net/http"
	"strings"
	"userAuthority/api/maps"
	"userAuthority/api/models"
	"userAuthority/api/thirdUtils"
	"userAuthority/api/validates"
)

// 解决跨域
func Cors(t *sunrise.Context) {
	method := t.Request.Method
	t.Header("Access-Control-Allow-Origin", "*")
	t.Header("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	t.Header("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	t.Header("Content-Type", "application/json;charset=utf-8")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		t.AbortWithStatus(http.StatusNoContent)
		return
	}
	t.Next()
}

//JWT以及用户权限验证
func ValidateJwtMiddleware(t *sunrise.Context) {
	const (
		SecretKey = "sunrise"
	)
	var url string
	var keyUserID string
	token, err := request.ParseFromRequest(t.Request, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	if err == nil {
		if token.Valid {
			userInfo, _ := jwtInfo.JwtParseUser(token.Raw)
			models.GetModelsEngine()
			keyUserID = thirdUtils.RedisGet(userInfo.UserID)
			user := models.GetUserInfoByUserID(userInfo.UserID)
			fmt.Println(len(user.UserID))
			if len(user.UserID) <= 0 && len(keyUserID) <= 0 {
				resp := maps.ResponseInfo{
					Code: 10,
					Msg:  "Token 验证用户失败！",
				}
				t.AbortWithStatusJSON(http.StatusOK, resp)
				return
			}
			webUrl := t.GetHeader("url")
			if len(webUrl) > 0 {
				url = webUrl
			} else {
				url = t.Request.RequestURI
			}
			isTrue := models.GetFunctionByUserID(userInfo.UserID, url)
			if !isTrue {
				resp := maps.ResponseInfo{
					Code:   8,
					Msg:    "用户权限不足,请联系管理员！",
					Result: nil,
				}
				t.AbortWithStatusJSON(http.StatusOK, resp)
				return
			}
			if len(webUrl) > 0 && isTrue {
				resp := maps.ResponseInfo{
					Code:   0,
					Msg:    "用户权限验证成功！",
					Result: user,
				}
				t.AbortWithStatusJSON(http.StatusOK, resp)
				return
			}
			t.Set("contextJwtUserInfo", *userInfo)
			t.Next()
		} else {
			resp := maps.ResponseInfo{
				Code: 10,
				Msg:  "token验证失败",
			}
			t.AbortWithStatusJSON(http.StatusOK, resp)
			return
		}
	} else {
		resp := maps.ResponseInfo{
			Code: 10,
			Msg:  "token验证失败",
		}
		t.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}
}

// 参数层的验证
func CheckValidate(t *sunrise.Context) {
	HandlerPath := t.HandlerName()           //包.方法体的格式   比如controllers下面的user.Login方法
	index := strings.Index(HandlerPath, ".") //“.”可以取出具体的方法体
	var msg string
	HandlerName := HandlerPath[index+1:]
	var resp maps.ResponseInfo
	err := models.GetModelsEngine()
	if err != nil {
		resp.Code = 90
		resp.Msg = maps.Msg[resp.Code]
		t.AbortWithStatusJSON(http.StatusOK, resp)
		return
	}

	var checkState bool
	switch HandlerName {
	case "WeChatAppletLogin": //登录
		var WeChatAppletLogin maps.WeChatAppletLoginRequest
		t.BindJSON(&WeChatAppletLogin)
		checkState, msg = validates.CheckParameter(WeChatAppletLogin)
		t.Set("WeChatAppletLogin", WeChatAppletLogin)
	case "UpdatePwd": //修改密码
		var UpdatePwd maps.UpdatePwd
		t.BindJSON(&UpdatePwd)
		checkState, msg = validates.CheckParameter(UpdatePwd)
		t.Set("UpdatePwd", UpdatePwd)
	case "InsertUser": //创建用户
		var InsertUser maps.InsertUser
		t.BindJSON(&InsertUser)
		checkState, msg = validates.CheckParameter(InsertUser)
		t.Set("InsertUser", InsertUser)
	case "UpdateUser": //修改用户
		var UpdateUser maps.UpdateUser
		t.BindJSON(&UpdateUser)
		checkState, msg = validates.CheckParameter(UpdateUser)
		t.Set("UpdateUser", UpdateUser)
	case "GetUserListInfo": //获取用户列表
		var GetUserListInfo maps.GetUserListInfo
		t.BindJSON(&GetUserListInfo)
		checkState, msg = validates.CheckParameter(GetUserListInfo)
		t.Set("GetUserListInfo", GetUserListInfo)
	case "UpdateUserIsDelete": //删除用户
		var UpdateUserIsDelete maps.UpdateUserIsDelete
		t.BindJSON(&UpdateUserIsDelete)
		checkState, msg = validates.CheckParameter(UpdateUserIsDelete)
		t.Set("UpdateUserIsDelete", UpdateUserIsDelete)
	case "InsertRole": //新增角色
		var InsertRole models.S_role
		t.BindJSON(&InsertRole)
		checkState, msg = validates.CheckParameter(InsertRole)
		t.Set("InsertRole", InsertRole)
	case "UpdateRole": //修改角色
		var UpdateRole models.S_role
		t.BindJSON(&UpdateRole)
		checkState, msg = validates.CheckParameter(UpdateRole)
		t.Set("UpdateRole", UpdateRole)
	case "DeleteRoleByRoleID": //删除角色
		var DeleteRoleByRoleID maps.DeleteRoleByRoleID
		t.BindJSON(&DeleteRoleByRoleID)
		checkState, msg = validates.CheckParameter(DeleteRoleByRoleID)
		t.Set("DeleteRoleByRoleID", DeleteRoleByRoleID)
	case "GetRoleFunctionByRoleID": //获取角色权限列表
		var GetRoleFunctionByRoleID maps.GetRoleFunctionByRoleIDRequest
		t.BindJSON(&GetRoleFunctionByRoleID)
		checkState, msg = validates.CheckParameter(GetRoleFunctionByRoleID)
		t.Set("GetRoleFunctionByRoleID", GetRoleFunctionByRoleID)
	case "InsertRoleFunction": //添加角色权限
		var InsertRoleFunction maps.InsertRoleFunctionInfo
		t.BindJSON(&InsertRoleFunction)
		checkState, msg = validates.CheckParameter(InsertRoleFunction)
		t.Set("InsertRoleFunction", InsertRoleFunction)
	default:
		checkState = true
	}
	if !checkState {
		resp := maps.ResponseInfo{
			Code: 1,
			Msg:  "参数错误" + msg,
		}
		t.AbortWithStatusJSON(http.StatusOK, resp)
		return
	} else {
		t.Next()
		return
	}
}
