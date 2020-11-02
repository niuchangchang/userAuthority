/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-08-28 17:57:38
 * @LastEditors: Dawn
 */
package models

import (
	"encoding/json"
	"fmt"
	"time"
	"userAuthority/api/maps"
	"userAuthority/api/thirdUtils"
)

type S_user_role struct {
	UrID       string    `json:"urID"  xorm:"notnull 'ID' default()"`         // 用户权限标识
	UserID     string    `json:"userID"  xorm:"notnull 'userID' default('')"` // 用户ID
	RoleID     string    `json:"roleID"  xorm:"notnull 'roleID' default('')"` // 角色ID
	InsertTime time.Time `json:"insertTime" xorm:"created 'insertTime'"`      // 记录插入时间
	UpdateTime time.Time `json:"updateTime" xorm:"updated 'updateTime'"`      // 记录更新时间
}

func (user_role S_user_role) TableName() string {
	return "s_user_role"
}

// 根据userID获取角色详情
func GetRoleIDByUserID(userID string) (roleInfo S_user_role) {
	exist, err := Engine.Where("userID = ?", userID).Get(&roleInfo)
	if exist == false || err != nil {
		fmt.Println(err, "GetRoleIDByUserID 失败原因")
		return S_user_role{}
	}
	return roleInfo
}

// 新增用户角色关系
func InsertUserRole(insertUserRole maps.UserRoleInfo) (msg string) {
	uuid := thirdUtils.UUID()
	insertUserRole.ID = uuid
	userRoleInsertInfo := new(S_user_role)
	b, _ := json.Marshal(insertUserRole)
	json.Unmarshal(b, &userRoleInsertInfo)
	affected, err := Engine.Insert(userRoleInsertInfo)
	if err != nil {
		fmt.Println(err, "InsertUserRole失败原因")
		return "新增失败！"
	} else {
		if affected > 0 {
			return "新增成功！"
		} else {
			return "新增失败！"
		}
	}
}
func DelUserRoleByUserID(UserID string) (msg string) {
	if exist, err := Engine.Exec("DELETE FROM s_user_role WHERE userID = ?", UserID); err != nil {
		fmt.Println(err, "DelUserRoleByUserID失败原因")
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

// 根据userID账号获取角色
func GetUserRoleByUserID(userID string) (roleInfo []map[string]string, err error) {
	sqlStr := `select * from s_user_role where userID=?`
	exist, err := Engine.QueryString(sqlStr, userID)
	if err != nil {
		fmt.Println(err, "GetUserRoleByUserID失败原因")
		return exist, err
	}
	return exist, nil
}
