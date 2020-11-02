/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-09-24 15:33:15
 * @LastEditors: Dawn
 */
package models

import (
	"fmt"
	"time"
	"userAuthority/api/thirdUtils"
)

type S_role struct {
	RoleID      string    `json:"roleID"  xorm:"notnull 'roleID' default('')"` // 角色标识
	RoleName    string    `json:"roleName"  xorm:"'roleName' default('')"`     // 角色名称
	OrderNum    int64     `json:"orderNum"  xorm:"'orderNum' default('')"`     // 角色顺序
	Description string    `json:"description"  xorm:"'description' default()"` // 角色描述
	InsertTime  string    `json:"insertTime" xorm:"created 'insertTime'"`      // 记录插入时间
	UpdateTime  time.Time `json:"updateTime" xorm:"updated 'updateTime'"`      // 记录更新时间
}

func (role S_role) TableName() string {
	return "s_role"
}

// 查询所有角色
func GetRoleList() (roleList []S_role, err error) {
	var roleArr []S_role
	if err := Engine.Find(&roleList); err != nil {
		fmt.Println(err, "GetRoleList失败原因")
		return roleList, err
	}

	for _, item := range roleList {

		roleArr = append(roleArr, item)

	}

	return roleArr, nil
}

// 新增角色
func InsertRole(insertRole S_role) (msg string) {
	uuid := thirdUtils.UUID()
	insertRole.RoleID = uuid
	// roleInsertInfo := new(S_role)
	// b, _ := json.Marshal(insertRole)
	// json.Unmarshal(b, &userInsertInfo)
	affected, err := Engine.Insert(insertRole)
	if err != nil {
		fmt.Println(err, "InsertRole失败原因")
		return "新增失败！"
	} else {
		if affected > 0 {
			return "新增成功！"
		} else {
			return "新增失败！"
		}
	}
}

// 修改角色
func UpdateRole(updateRole S_role) (msg string) {
	// userUpdateInfo := new(S_user)
	// b, _ := json.Marshal(updateRole)
	// json.Unmarshal(b, &userUpdateInfo)
	// userUpdateInfo.LockTime = time.Now()
	affected, err := Engine.Where("roleID=?", updateRole.RoleID).Update(updateRole)
	if err != nil {
		fmt.Println(err, "UpdateRole失败原因")
		return "修改失败！"
	} else {
		if affected > 0 {
			return "修改成功！"
		} else {
			return "修改失败！"
		}
	}
}

// 根据角色名称获取用户详情
func GetRoleInfoByRoleName(roleName string) (hs bool, err error) {
	has, err := Engine.SQL("select * from s_role where roleName = ?", roleName).Exist()
	if err != nil {
		fmt.Println(err, "GetRoleInfoByRoleName失败原因")
		return has, err
	}
	return has, nil
}

// 根据角色ID删除
func DeleteRoleByRoleID(roleID string) (msg string) {
	sql := "DELETE FROM s_role where roleID=?"
	exist, err := Engine.Exec(sql, roleID)
	if err != nil {
		fmt.Println(err, "GetRoleInfoByRoleName失败原因")
		return "删除失败！"
	}
	info, _ := exist.RowsAffected()
	if info > 0 {
		return "删除成功！"
	} else {
		return "删除失败！"
	}
}
