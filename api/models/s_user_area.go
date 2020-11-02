/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-07 11:03:57
 * @LastEditTime: 2019-09-23 11:32:08
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

type S_user_area struct {
	UaID        *string   `json:"uaID"  xorm:"notnull 'uaID' default('')"`               // 权限区域关系标识
	VillageID   *string   `json:"villageID"  xorm:"notnull 'villageID' default('')"`     // 小区标识
	UserID      *string   `json:"userID"  xorm:"notnull 'userID' default(0)"`            // 用户标识
	AreaCode    *string   `json:"areaCode"  xorm:"notnull 'areaCode' default('')"`       // 行政区域标识根据区域Code来关联：省-城市-行政区-街道-小区
	AreaName    *string   `json:"areaName"  xorm:"notnull 'areaName' default('')"`       // 行政区域中文 eg:上海-上海市-徐汇区-田林街道
	VillageName *string   `json:"villageName"  xorm:"notnull 'villageName' default('')"` // 行政区域中文 eg:上海-上海市-徐汇区-田林街道
	InsertTime  time.Time `json:"insertTime" xorm:"created 'insertTime'"`                // 记录插入时间
	UpdateTime  time.Time `json:"updateTime" xorm:"updated 'updateTime'"`                // 记录更新时间

}

func (user_area S_user_area) TableName() string {
	return "s_user_area"
}

// 新增用户区域
func InsertUserArea(userAreaInfo maps.UserAreaInfo) (msg string) {
	userAreaInsertInfo := new(S_user_area)
	b, _ := json.Marshal(userAreaInfo)
	json.Unmarshal(b, &userAreaInsertInfo)
	uuid := thirdUtils.UUID()
	userAreaInsertInfo.UaID = &uuid
	affected, err := Engine.Insert(userAreaInsertInfo)
	if err != nil {
		fmt.Println(err, "InsertUserArea失败原因")
		return "新增失败！"
	} else {
		if affected > 0 {
			return "新增成功！"
		} else {
			return "新增失败！"
		}
	}
}
func DelUserAreaByUserID(UserID string) (msg string) {
	if exist, err := Engine.Exec("DELETE FROM s_user_area WHERE userID = ?", UserID); err != nil {
		fmt.Println(err, "DelUserAreaByUserID失败原因")
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

// 通过userID 获取用户对应 areaCode 列表
func GetAreaCodeListByUserID(userID string) (userAreaList []S_user_area, err error) {
	err = Engine.Cols("areaCode", "areaName").Where("userID = ?", userID).Find(&userAreaList)
	if err != nil {
		fmt.Println(err, "GetAreaCodeListByUserID失败原因")
		return userAreaList, err
	}

	return userAreaList, nil

}
