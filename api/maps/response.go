// Package maps ...响应参数
/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-05 10:28:04
 * @LastEditTime: 2019-10-23 15:25:05
 * @LastEditors: Dawn
 */
package maps

import (
	"time"
)

type ResponseInfo struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Result interface{} `json:"result,omitempty"`
}
type S_userInfo struct {
	UserID      string    `json:"userID"`      // 系统用户标识
	LoginName   string    `json:"loginName"`   // 系统用户编号
	DisplayName string    `json:"displayName"` // 显示名称
	PhoneTel    string    `json:"phoneTel"`    // 手机号码
	Token       string    `json:"token"`       // token
	Status      int64     `json:"status"`      // 账号状态 0-正常  1-禁用
	IsDelete    int64     `json:"isDelete"`    // 是否有效： 0-正常  1-禁用
	IsSysOrApp  int64     `json:"isSysOrApp"`  // 0-系统用户1-APP用户
	LockTime    time.Time `json:"lockTime"`    // 锁定时间
	LoginNum    int64     `json:"loginNum"`    // 登录失败次数
	IsLock      int64     `json:"isLock"`      // 登录方式：0-无锁 1-锁定
	// Token        string    `json:"token"         `                                             // 登录方式：0-无锁 1-锁定
	HeadPicUrl string    ` json:"headPicUrl"`  // 头像
	InsertTime time.Time ` json:"insertTime "` // 记录插入时间
	UpdateTime time.Time ` json:"updateTime "` // 记录更新时间

	RoleID []map[string]string `json:"roleID"` //角色

	// ProvinceLists  []map[string]string `json:"provinceLists"`  //区域
	// CityLists      []map[string]string `json:"cityLists"`      //区域
	// DistrictLists  []map[string]string `json:"districtLists"`  //区域
	// StreetLists    []map[string]string `json:"streetLists"`    //区域
	CommitteeLists []map[string]string `json:"committeeLists"` //区域
	VillageLists   []map[string]string `json:"villageLists"`   //区域
}

// type FunctionList struct {
// 	FunctionList []map[string]string `json:"functionList"`
// }
type FunctionInfo struct {
	FunctionID         string         `json:"functionID"  `         // 主键
	ParentFunctionCode string         `json:"parentFunctionCode"  ` // 功能编码父节点
	FunctionCode       string         `json:"functionCode"        ` // 功能编码
	FunctionName       string         `json:"functionName"        ` // 功能名称
	FunctionType       int64          `json:"functionType"        ` // 功能类型(0菜单 1按钮)
	FunctionMethod     string         `json:"functionMethod"      ` // 请求方法:get、post、delete、put
	FunctionIcon       string         `json:"functionIcon"        ` // 菜单图标
	Description        string         `json:"description"         ` // 功能描述
	IsValid            int64          `json:"isValid"             ` // 是否有效 0-有效 1-无效 0-有效  1-无效
	Path               string         `json:"path"                `
	Children           []FunctionInfo `json:"children"                `
}
type UserInfoList struct {
	UserID      string `json:"userID"`      // 系统用户标识
	LoginName   string `json:"loginName"`   // 系统用户编号
	DisplayName string `json:"displayName"` // 显示名称
	Email       string `json:"email"      ` // 电子邮箱

	PhoneTel string `json:"phoneTel"` // 手机号码
	Status   int64  `json:"status"`   // 账号状态 0-正常  1-禁用
	// IsDelete   int64     `json:"isDelete"`   // 是否有效： 0-正常  1-禁用
	IsSysOrApp int64     `json:"isSysOrApp"` // 0-系统用户1-APP用户
	LockTime   time.Time `json:"lockTime"`   // 锁定时间
	IsLock     int64     `json:"isLock"`     // 登录方式：0-无锁 1-锁定
	// Token        string    `json:"token"         `                                             // 登录方式：0-无锁 1-锁定
	HeadPicUrl string ` json:"headPicUrl"` // 头像
	InsertTime string ` json:"insertTime"` // 记录插入时间
	Count      string `json:"count"`
	// UserAreaResponseInfo []interface{} `json:"userAreaResponseInfo"`
	AreaNameResponseInfo    []map[string]string `json:"userAreaResponseInfo"`
	AreaResponseInfo        []map[string]string `json:"areaResponseInfo"`
	VillageNameResponseInfo []map[string]string `json:"villageNameResponseInfo"`

	RoleID []map[string]string `json:"roleID"`
}

type UserAreaResponseInfo struct {
	VillageID string `json:"villageID" `
	AreaCode  string `json:"areaCode" validate:"required"` // 行政区域标识根据区域Code来关联：省-城市-行政区-街道-小区
	AreaName  string `json:"areaName" validate:"required"` // 行政区域中文 eg:上海-上海市-徐汇区-田林街道
}

type GetUserListResponseInfo struct {
	UserList []UserInfoList `json:"userList" `
	Count    int64          `json:"Count" ` // 行政区域标识根据区域Code来关联：省-城市-行政区-街道-小区
}
type GetDeviceListResponseInfo struct {
	DeviceList []map[string]string `json:"deviceList" `
	Count      int64               `json:"count" `
}
type BatchInsertDeviceResponseInfo struct {
	ResultIs   []string `json:"resultIs" `
	ResultFull []string `json:"resultFull" `
	ResultIp   []string `json:"resultIp" `
	ResultProject []string `json:"resultProject" `
}
type GetSystemLogListResponseInfo struct {
	SysList []map[string]string `json:"sysList" `
	Count   int64               `json:"Count" ` // 行政区域标识根据区域Code来关联：省-城市-行政区-街道-小区
}
type Authorization struct {
	Authorization string `json:"authorization" `
}

