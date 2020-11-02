/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-10-25 16:22:10
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

var cstZone = time.FixedZone("CST", 8*3600)

type S_user struct {
	UserID      string    `json:"userID"  xorm:"notnull 'userID' default('')"` // 系统用户标识
	LoginName   *string   `json:"loginName"  xorm:"'loginName' default()"`     // 系统用户编号
	Password    *string   `json:"password"  xorm:"'password' default()"`       // 登录密码通过加密的二进制保存
	DisplayName *string   `json:"displayName"  xorm:"'displayName' default()"` // 显示名称
	Gender      *string   `json:"gender"  xorm:"'gender' default()"`           // 性别
	BirthDate   *string   `json:"birthDate"  xorm:"'birthDate' default()"`     // 生日
	PhoneTel    *string   `json:"phoneTel"  xorm:"'phoneTel' default()"`       // 手机号码
	Tel         *string   `json:"tel"  xorm:"'tel' default()"`                 // 座机号码
	Position    *string   `json:"position"  xorm:"'position' default()"`       // 职位名字
	Email       *string   `json:"email"  xorm:"'email' default()"`             // 电子邮箱
	Status      *int64    `json:"status"  xorm:"'status' default(0)"`          // 账号状态 0-正常  1-禁用
	IsDelete    int64     `json:"isDelete"  xorm:"'isDelete' default(0)"`      // 是否有效： 0-正常  1-禁用
	HeadPicUrl  *string   `json:"headPicUrl"  xorm:"'headPicUrl' default()"`   // 头像
	InsertTime  time.Time `json:"insertTime" xorm:"created 'insertTime'"`      // 记录插入时间
	UpdateTime  time.Time `json:"updateTime" xorm:"updated 'updateTime'"`      // 记录更新时间
}

func (user S_user) TableName() string {
	return "s_user"
}

//// 用户登录获取用户信息
//func GetUserInfoByLogin(login maps.Login) (userInfo S_user, err error) {
//	exist, err := Engine.Where("loginName = ? AND password = ? and isSysOrApp=0 and status =0 and isLock=0  AND isDelete=0 ", login.LoginName, login.Password).Get(&userInfo)
//	if exist == false || err != nil {
//		fmt.Println(err, "GetUserInfoByLogin失败原因")
//		return S_user{}, errors.New("查询用户信息失败")
//	}
//	return userInfo, nil
//}

// 根据登录账号获取用户详情
func GetUserInfoByLoginName(loginname string) (userInfo S_user) {
	exist, err := Engine.Where("loginName = ? AND isDelete=1 AND status=1 ", loginname).Get(&userInfo)
	if exist == false || err != nil {
		fmt.Println(err, "GetUserInfoByLoginName 失败原因")
		return S_user{}
	}
	return userInfo
}

// 根据userID获取用户详情
func GetUserInfoByUserID(userID string) (userInfo S_user) {
	exist, err := Engine.Where("userID = ? AND isDelete=1 ", userID).Get(&userInfo)
	if exist == false || err != nil {
		fmt.Println(err, "GetUserInfoByUserID 失败原因")
		return S_user{}
	}
	return userInfo
}

// 根据登录账号获取用户详情 （手机用户）
func GetUserInfoByTel(loginname string) (userInfo S_user, err error) {
	exist, err := Engine.Where("loginName = ? AND isSysOrApp=1 AND isDelete=0 ", loginname).Get(&userInfo)
	if exist == false || err != nil {
		fmt.Println(err, "GetUserInfoByTel失败原因")
		return S_user{}, err
	}
	return userInfo, nil
}

// 修改密码
func UpdatePwd(updatePwd maps.UpdatePwd) (msg string, code int) {
	if exist, err := Engine.Exec("update s_user set password = ? where userID = ? and password=? ", updatePwd.Password, updatePwd.UserID, updatePwd.OldPassWord); err != nil {
		fmt.Println(err, "UpdatePwd失败原因")
		return "修改失败！", 92
	} else {
		info, _ := exist.RowsAffected()
		if info > 0 {
			return "修改成功！", 0
		} else {
			return "旧密码密码错误或修改失败！", 92
		}
	}
}

// 新增用户
func InsertUser(insertUser maps.InsertUser) (msg string, userID string) {
	uuid := thirdUtils.UUID()
	insertUser.UserID = uuid
	userInsertInfo := new(S_user)
	b, _ := json.Marshal(insertUser)
	json.Unmarshal(b, &userInsertInfo)
	//is := int64(0)
	//userInsertInfo.LoginNum = &is
	affected, err := Engine.Insert(userInsertInfo)
	if err != nil {
		fmt.Println(err, "InsertUser失败原因")

		return "新增失败！", ""
	} else {
		if affected > 0 {
			return "新增成功！", insertUser.UserID
		} else {
			return "新增失败！", ""
		}
	}
}

// 修改用户
func UpdateUser(updateUser maps.UpdateUser) (msg string) {
	userUpdateInfo := new(S_user)
	b, _ := json.Marshal(updateUser)
	json.Unmarshal(b, &userUpdateInfo)
	if *(updateUser.IsLock) == 0 {
		//is := int64(0)
		//userUpdateInfo.LoginNum = &is
	}
	affected, err := Engine.Where("userID=?", userUpdateInfo.UserID).Update(userUpdateInfo)
	if err != nil {
		fmt.Println(err, "UpdateUser失败原因")
		return "修改失败！"
	} else {
		if affected > 0 {
			return "修改成功！"
		} else {
			return "修改失败！"
		}
	}
}

// 修改用户登录失败次数
func UpdateUserNum(updateUserNum maps.UpdateNum) (msg string) {
	if exist, err := Engine.Exec("update s_user set loginNum = ? where loginName = ? and isSysOrApp=0", updateUserNum.LoginNum, updateUserNum.Loginname); err != nil {
		fmt.Println(err, "UpdateUserNum失败原因")
		return "修改失败！"
	} else {
		info, _ := exist.RowsAffected()
		fmt.Println(info, "修改登录次数状态")
		if info > 0 {
			return "修改成功！"
		} else {
			return "修改失败！"
		}
	}
}

// 修改用户锁定状态
func UpdateUserIsLock(Loginname string, isLock int64) {
	Engine.Exec("update s_user set isLock = ?, lockTime=? where loginName = ? and isSysOrApp=0", isLock, time.Now(), Loginname)
}

// 伪删除用户
func UpdateUserIsDelete(userID string) {
	Engine.Exec("update s_user set isDelete = 1 where userID = ? ", userID)
}

func GetUserListInfoCount(userInfo maps.GetUserListInfo) (count int64) {
	var userList []map[string]string
	var vill string
	for i := 0; i < len(userInfo.VillageIDs); i++ {
		if i == len(userInfo.VillageIDs)-1 {
			vill = vill + "'" + userInfo.VillageIDs[i] + "'"
		} else {
			vill = vill + "'" + userInfo.VillageIDs[i] + "'" + ","
		}
	}
	fmt.Println(vill, "<----------------")
	sqlStr := `select s_user.* from s_user LEFT JOIN s_user_area on s_user_area.userID=s_user.userID where 1=1 and s_user.isDelete=0`
	if len(userInfo.DisplayName) > 0 {
		sqlStr += ` and s_user.displayName=` + "'" + userInfo.DisplayName + "'"
	}
	if len(userInfo.VillageIDs) > 0 {
		sqlStr += ` and s_user_area.villageID in( ` + vill + `)`
	}
	if len(userInfo.LoginName) > 0 {
		sqlStr += ` and s_user.loginName=` + "'" + userInfo.LoginName + "'"
	}
	if len(userInfo.PhoneTel) > 0 {
		sqlStr += ` and s_user.phoneTel= ` + "'" + userInfo.PhoneTel + "'"
	}
	if *userInfo.IsSysOrApp == 0 {
		IsSysOrApp := "0"
		sqlStr += ` and s_user.isSysOrApp= ` + IsSysOrApp
	}
	if *userInfo.IsSysOrApp == 1 {
		IsSysOrApp := "1"
		sqlStr += ` and s_user.isSysOrApp= ` + IsSysOrApp
	}
	sqlStr += ` GROUP BY s_user.userID `
	userList, _ = Engine.QueryString(sqlStr)
	lens := len(userList)
	return int64(lens)

}
func GetUserListInfo(userInfo maps.GetUserListInfo) (userInfoList []maps.UserInfoList) {
	var userList []map[string]string
	var vill string
	for i := 0; i < len(userInfo.VillageIDs); i++ {
		if i == len(userInfo.VillageIDs)-1 {
			vill = vill + "'" + userInfo.VillageIDs[i] + "'"
		} else {
			vill = vill + "'" + userInfo.VillageIDs[i] + "'" + ","
		}
	}
	fmt.Println(vill, "<----------------")
	sqlStr := `select s_user.userID,s_user.loginName,s_user.displayName,s_user.phoneTel,s_user.status,s_user.isSysOrApp,s_user.isLock,s_user.insertTime,s_user.email from s_user LEFT JOIN s_user_area on s_user_area.userID=s_user.userID where 1=1 and s_user.isDelete=0 `
	startIndex := (*userInfo.PageNum - 1) * *userInfo.PageSize
	pageSize := *userInfo.PageSize
	if len(userInfo.DisplayName) > 0 {
		sqlStr += ` and s_user.displayName=` + "'" + userInfo.DisplayName + "'"
	}
	if len(userInfo.VillageIDs) > 0 {
		sqlStr += ` and s_user_area.villageID in( ` + vill + `)`
	}
	if len(userInfo.LoginName) > 0 {
		sqlStr += ` and s_user.loginName=` + "'" + userInfo.LoginName + "'"
	}
	if len(userInfo.PhoneTel) > 0 {
		sqlStr += ` and s_user.phoneTel= ` + "'" + userInfo.PhoneTel + "'"
	}
	if *userInfo.IsSysOrApp == 0 {
		IsSysOrApp := "0"
		sqlStr += ` and s_user.isSysOrApp= ` + IsSysOrApp
	}
	if *userInfo.IsSysOrApp == 1 {
		IsSysOrApp := "1"
		sqlStr += ` and s_user.isSysOrApp= ` + IsSysOrApp
	}
	sqlStr += ` GROUP BY s_user.userID ORDER BY s_user.updateTime limit ?,? `
	userList, _ = Engine.QueryString(sqlStr, startIndex, pageSize)
	for _, value := range userList {
		var user maps.UserInfoList
		user.UserID = value["userID"]
		user.Email = value["email"]
		user.DisplayName = value["displayName"]
		user.Status, _ = strconv.ParseInt(value["status"], 10, 64)
		user.PhoneTel = value["phoneTel"]
		user.LoginName = value["loginName"]
		user.IsSysOrApp, _ = strconv.ParseInt(value["isSysOrApp"], 10, 64)
		user.IsLock, _ = strconv.ParseInt(value["isLock"], 10, 64)
		user.UserID = value["userID"]
		user.InsertTime, _ = value["insertTime"]
		sql := `SELECT GROUP_CONCAT(s_user_area.areaName) as areaName ,GROUP_CONCAT(s_user_area.areaCode) as areaCode from s_user_area where userID=? `
		areaName, _ := Engine.QueryString(sql, user.UserID)
		user.AreaNameResponseInfo = areaName

		sqlarea := `SELECT GROUP_CONCAT(s_user_area.areaName) as areaName ,GROUP_CONCAT(s_user_area.areaCode) as areaCode, GROUP_CONCAT(case when s_user_area.villageName<>' ' THEN s_user_area.villageName else null END) as villageName,GROUP_CONCAT( case when s_user_area.villageID<>'' THEN s_user_area.villageID else null END ) as villageID from s_user_area where userID=? `
		area, _ := Engine.QueryString(sqlarea, user.UserID)
		user.AreaResponseInfo = area

		sqlV := `SELECT GROUP_CONCAT(case when s_user_area.villageName<>' ' THEN s_user_area.villageName else null END) as villageName,GROUP_CONCAT( case when s_user_area.villageID<>'' THEN s_user_area.villageID else null END ) as villageID from s_user_area where userID=? `
		villageName, _ := Engine.QueryString(sqlV, user.UserID)
		if len(villageName) <= 0 {
			user.VillageNameResponseInfo = nil
		}
		user.VillageNameResponseInfo = villageName
		roleIDsql := `SELECT s_role.roleID,s_role.roleName from s_user_role LEFT JOIN s_role on s_user_role.roleID=s_role.roleID
    where userID=? `
		roleID, _ := Engine.QueryString(roleIDsql, user.UserID)
		user.RoleID = roleID
		userInfoList = append(userInfoList, user)
	}
	return userInfoList

}
