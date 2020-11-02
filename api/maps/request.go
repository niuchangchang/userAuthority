/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-14 11:25:58
 * @LastEditTime: 2019-10-23 15:16:52
 * @LastEditors: Dawn
 */
// Package maps ...请求参数

package maps

// WebLogin makes maps.
type WebLogin struct {
	LoginName string `json:"loginName"` //账号
	Password  string `json:"password" ` //密码
}

// AndroidLogin makes maps.
type AndroidLogin struct {
	LoginName string `json:"loginName"` //账号
	SmsCode   string `json:"smsCode" `  //短信验证码
}

// WeChatAppletLoginRequest makes maps.
type WeChatAppletLoginRequest struct {
	WeChatAppletType string `json:"weChatAppletType"  validate:"required"` // 1-驾驶端 、企业端 2-销售端
	Code             string `json:"code"  validate:"required" `            //微信授权code
	EncryptedData    string `json:"encryptedData"  validate:"required"`    //微信授权encryptedData
	Iv               string `json:"iv"  validate:"required" `              //微信授权iv
}

type GetRoleFunctionByRoleIDRequest struct {
	Platform string `json:"platform"  validate:"required"` // 1-驾驶端 、企业端 2-销售端
}
type JwtUserInfo struct {
	UserID   string `json:"userID"`
	Phone    string `json:"phone"`
	OpenID   string `json:"openID"`
	UserName string `json:"userName"`
	Address  string `json:"address"`
	Remark   string `json:"remark"`
}

type UpdateNum struct {
	LoginNum  int64  `json:"loginNum" validate:"required"`  //登录失败次数
	Loginname string `json:"loginName" validate:"required"` //用户名

}
type UpdatePwd struct {
	UserID      string `json:"userID" validate:"required"`      //用户ID
	Password    string `json:"password" validate:"required"`    //密码
	OldPassWord string `json:"oldPassWord" validate:"required"` //旧密码
}

type GetCityListByProvinceCode struct { //获取市
	ProvinceCode string `json:"provinceCode" validate:"required"` //省code
}
type GetDistrictListByCityCode struct { //获取区
	CityCode string `json:"cityCode" validate:"required"` //市code
}
type GetStreetListByDistrictCode struct { //获取街道
	DistrictCode string `json:"districtCode" validate:"required"` //区code
}
type GetCommitteeListByStreetCode struct { //获取居委
	StreetCode string `json:"streetCode" validate:"required"` //街道code
}
type GetVillageListByCommitteeCode struct { //获取小区
	CommitteeCode string `json:"committeeCode" validate:"required"` //居委code
}

type InsertUser struct {
	UserID           string         `json:"userID"   `                        // 系统用户标识
	LoginName        string         `json:"loginName"  validate:"required" `  // 系统用户编号
	Password         string         `json:"password"   validate:"required" `  // 登录密码通过加密的二进制保存
	DisplayName      string         `json:"displayName" validate:"required" ` // 显示名称
	PhoneTel         string         `json:"phoneTel"   validate:"required" `  // 手机号码
	Email            string         `json:"email"       `                     // 电子邮箱
	Status           *int64         `json:"status"     validate:"required" `  // 账号状态 0-正常  1-禁用
	IsSysOrApp       *int64         `json:"isSysOrApp" validate:"required" `  // 0-系统用户1-APP用户
	IsLock           *int64         `json:"isLock"     validate:"required" `  // 登录方式：0-无锁 1-锁定
	UserRoleInfoList []UserRoleInfo `json:"userRoleInfoList" `
	UserAreaInfoList []UserAreaInfo `json:"userAreaInfoList"  validate:"required"`
}
type UserAreaInfo struct {
	UaID        string `json:"uaID"`                         // 权限区域关系标识
	UserID      string `json:"userID"   `                    // 用户标识
	AreaCode    string `json:"areaCode" validate:"required"` // 行政区域标识根据区域Code来关联：省-城市-行政区-街道-小区
	AreaName    string `json:"areaName" validate:"required"` // 行政区域中文 eg:上海-上海市-徐汇区-田林街道
	VillageName string `json:"villageName"`                  // 小区名称
	VillageID   string `json:"villageID"`                    // 小区ID

}
type UpdateUser struct {
	UserID           string         `json:"userID"     validate:"required" `  // 系统用户标识
	DisplayName      string         `json:"displayName" validate:"required" ` // 显示名称
	PhoneTel         string         `json:"phoneTel"   validate:"required" `  // 手机号码
	Email            string         `json:"email"      `                      // 电子邮箱
	Status           *int64         `json:"status"     validate:"required" `  // 账号状态 0-正常  1-禁用
	IsLock           *int64         `json:"isLock"     validate:"required" `  // 登录方式：0-无锁 1-锁定
	LoginNum         *int64         `json:"loginNum"    `                     // 登录方式：0-无锁 1-锁定
	IsSysOrApp       *int64         `json:"isSysOrApp" validate:"required" `  // 0-系统用户1-APP用户
	UserRoleInfoList []UserRoleInfo `json:"UserRoleInfoList"  `
	UserAreaInfoList []UserAreaInfo `json:"UserAreaInfoList"  validate:"required"`
}

type DeleteRoleByRoleID struct {
	RoleID string `json:"roleID" validate:"required"` //角色ID
}

type UserRoleInfo struct {
	ID     string `json:"ID" validate:"required"`     // 用户权限标识
	UserID string `json:"userID" `                    // 用户ID
	RoleID string `json:"roleID" validate:"required"` // 角色ID
}

type RoleFunctionInfo struct {
	RoleID     string `json:"roleID" validate:"required"`     //用户角色
	RfID       string `json:"rfID" validate:"required"`       // 主键ID
	FunctionID string `json:"functionID" validate:"required"` // 权限ID
}
type InsertRoleFunctionInfo struct {
	RoleFunctionInfoList []RoleFunctionInfo `json:"roleFunctionInfoList" validate:"required"`
}
type UpdateUserIsDelete struct {
	UserID string `json:"userID" validate:"required"` // 用户ID
}

type GetUserListInfo struct {
	VillageIDs  []string `json:"villageIDs" `
	DisplayName string   `json:"displayName" `
	LoginName   string   `json:"loginName"`
	PhoneTel    string   `json:"phoneTel"` // 手机号码
	PageNum     *int64   `json:"pageNum" validate:"required"`
	PageSize    *int64   `json:"pageSize" validate:"required"`
	IsSysOrApp  *int64   `json:"isSysOrApp" validate:"required"` //0-系统用户1-APP用户
}
type GetDeviceListInfo struct {
	IsInitType   string   `json:"isInitType" validate:"required"` //1-工程管理2-配置管理
	VillageIDs   []string `json:"villageIDs" `                    //小区
	ProductModel string   `json:"productModel" `                  //产品型号
	Code         string   `json:"code"`                           //设备编号
	LocationCode string   `json:"locationCode"`                   // 位置编码
	PrducetBrand string   `json:"prducetBrand"`                   //设备厂商
	IP           string   `json:"iP"`                             //主机ip(字表)
	ProductSN    string   `json:"productSN"`                      //序列号
	InstallAddr  string   `json:"installAddr"`                    //安装地址
	InsertTime   string   `json:"insertTime"`                     //创建时间
	EndTime      string   `json:"endTime"`                        //截止时间
	InstallTime  string   `json:"installTime"`                    //安装时间
	DeviceType   *int64   `json:"deviceType" validate:"required"` //1-IPC 2-网关 3-门禁

	PageNum  *int64 `json:"pageNum" validate:"required"`
	PageSize *int64 `json:"pageSize" validate:"required"`
}
type GetSystemLogListInfo struct {
	PageNum  *int64 `json:"pageNum" validate:"required"`
	PageSize *int64 `json:"pageSize" validate:"required"`
}
type InsertDeviceInfo struct {
	DeviceType        string              `json:"deviceType" validate:"required"` //1-网关 门禁 2-摄像机
	InsertsDeviceInfo []InsertsDeviceInfo `json:"insertsDeviceInfo" validate:"required"`
	// AccessInfo []AccessInfo `json:"accessInfo"`
	// CameraInfo []CameraInfo `json:"cameraInfo"`
}
type InsertsDeviceInfo struct {
	//设备
	PrducetBrand *string `json:"prducetBrand" validate:"required"`  // 产品品牌
	ProductModel *string `json:"productModel" validate:"required" ` // 产品型号
	DeviceID     *string `json:"deviceID" validate:"required"`      // 设备ID
	LocationCode *string `json:"locationCode" validate:"required"`  // 位置编码/地址编码
	Code         *string `json:"code" validate:"required"`          // 设备编号
	AreaName     *string `json:"areaName" validate:"required"`      // 行政区域
	VillageID    *string `json:"villageID" validate:"required"`     // 小区标识
	InstallAddr  *string `json:"installAddr" `                      // 安装地址
	Type         *string `json:"type" validate:"required"`
	IsInit       *int64  `json:"isInit" validate:"required"`    //设备初始化状态:1-未初始化 2-初始化过
	IsDisable    *int64  `json:"isDisable" validate:"required"` //是否启用:0-禁用 1-启用
	State        *int64  `json:"state" validate:"required"`     //安装状态:0-离线 1-在线 2-故障
	//门禁，网关
	AccessID                   string  `json:"accessID" `                                      // AccessID
	SetNetworkWay              *int64  `json:"setNetworkWay" validate:"required"`              // 配网方式:0-手动配置 1-DCHP配置
	Ip                         *string `json:"ip" validate:"required"`                         // ip
	SubnetMask                 *string `json:"subnetMask" validate:"required" `                // 子网掩码
	GateWay                    *string `json:"gateWay" validate:"required"`                    // 网关
	RtspUrl                    *string `json:"rtspUrl" `                                       // 视频源流
	NTPServer                  *string `json:"NTPServer" validate:"required"`                  // NTP服务器
	NASServer                  *string `json:"NASServer" `                                     // NAS服务器
	DNSServer                  *string `json:"DNSServer" `                                     // DNS服务器
	OMServer                   *string `json:"OMServer" `                                      // OMS服务器
	MQTTServer                 *string `json:"MQTTServer" validate:"required"`                 // MQTT服务器
	SIPServer                  *string `json:"SIPServer" `                                     // SIPServer服务器
	BussinessServer            *string `json:"BussinessServer" validate:"required"`            // 物管
	IsOpenByBlueTooth          string  `json:"isOpenByBlueTooth" validate:"required"`          // 是否蓝牙开门
	IsOpenByTelenoun           string  `json:"isOpenByTelenoun" validate:"required"`           // 是否远程开门
	IsOpenByFlow               string  `json:"isOpenByFlow" validate:"required"`               // 是否一键开门
	IsShare                    string  `json:"isShare" validate:"required"`                    // 是否分享
	OpenMaxDistanceByBlueTooth string  `json:"openMaxDistanceByBlueTooth" validate:"required"` // 蓝牙开门最大距离
	OpenMaxDistanceByTelenoun  string  `json:"openMaxDistanceByTelenoun" validate:"required"`  // 远程开门最大距离
	OpenMaxDistanceByFlow      string  `json:"openMaxDistanceByFlow" validate:"required"`      // 一键开门最大距离
	//摄像机
	CameraWay          *int64  `json:"cameraWay" validate:"required"`         // 配网方式:0-手动配置 1-DCHP配置
	CameraID           string  `json:"cameraID" validate:"required"`          // ID
	CameraIP           *string `json:"cameraIP" validate:"required"`          // ip
	CameraSubnetMask   *string `json:"cameraSubnetMask" validate:"required" ` // 子网掩码
	FaceGateWay        *string `json:"faceGateWay" validate:"required"`       // 人脸网关
	CameraOMServer     *string `json:"cameraOMServer" `                       // OM服务器
	CameraNVRServer    *string `json:"cameraNVRServer" `                      // NVR服务器
	CameraNVRChannelNo *string `json:"cameraNVRChannelNo" `                   // 通道号 1
	InOutFlag          *int64  `json:"inOutFlag" validate:"required"`         // 安装方向0-门外或进1-门内或出2-广场或其他开阔性的地方
}

type UpdateDeviceInfo struct {
	DeviceType string `json:"deviceType" validate:"required"` //1-网关 门禁 2-摄像机
	//设备
	PrducetBrand *string `json:"prducetBrand" ` // 产品品牌 1
	ProductModel *string `json:"productModel" ` // 产品型号 1
	DeviceID     *string `json:"deviceID" `     // 设备ID
	AreaName     *string `json:"areaName" `     // 行政区域 1
	VillageID    *string `json:"villageID"`     // 小区标识 1
	InstallAddr  *string `json:"installAddr" `  // 安装地址 1
	IsDelete     *int64  `json:"isDelete" `     // 是否删除:0-否 1-是
	// IsInit       *int64  `json:"isInit" `       //设备初始化状态:1-未初始化 2-初始化过
	IsDisable *int64 `json:"isDisable" ` //是否启用:0-禁用 1-启用
	State     *int64 `json:"state" `     //安装状态:0-离线 1-在线 2-故障
	//门禁，网关
	AccessID                   *string `json:"accessID"`                   //ID
	SetNetworkWay              *int64  `json:"setNetworkWay" `             // 配网方式:0-手动配置 1-DCHP配置1
	Ip                         *string `json:"ip" `                        //ip 1
	SubnetMask                 *string `json:"subnetMask"  `               // 子网掩码 1
	GateWay                    *string `json:"gateWay" `                   // 网关 1
	RtspUrl                    *string `json:"rtspUrl" `                   // 视频源流1
	NTPServer                  *string `json:"NTPServer" `                 // NTP服务器1
	NASServer                  *string `json:"NASServer" `                 // NAS服务器1
	IMGServer                  *string `json:"IMGServer" `                 //IMG服务器1
	SIPServer                  *string `json:"SIPServer" `                 //SIP服务器1
	OMServer                   *string `json:"OMServer" `                  // OMS服务器 1
	MQTTServer                 *string `json:"MQTTServer" `                // MQTT服务器1
	DNSServer                  *string `json:"DNSServer" `                 // DNS服务器1
	BussinessServer            *string `json:"BussinessServer" `           // 物管 1
	IsOpenByBlueTooth          string  `json:"isOpenByBlueTooth"`          // 是否蓝牙开门
	IsOpenByTelenoun           string  `json:"isOpenByTelenoun"`           // 是否远程开门
	IsOpenByFlow               string  `json:"isOpenByFlow"`               // 是否一键开门
	IsShare                    string  `json:"isShare"`                    // 是否分享
	OpenMaxDistanceByBlueTooth string  `json:"openMaxDistanceByBlueTooth"` // 蓝牙开门最大距离
	OpenMaxDistanceByTelenoun  string  `json:"openMaxDistanceByTelenoun"`  // 远程开门最大距离
	OpenMaxDistanceByFlow      string  `json:"openMaxDistanceByFlow"`      // 一键开门最大距离
	//摄像机
	CameraID           *string `json:"cameraID" `           //ID
	CameraIP           *string `json:"cameraIP" `           //ip1
	CameraSubnetMask   *string `json:"cameraSubnetMask"  `  // 子网掩码 1
	FaceGateWay        *string `json:"faceGateWay" `        // 人脸网关1
	CameraOMServer     *string `json:"cameraOMServer" `     // OM服务器1
	CameraDNSServer    *string `json:"cameraDNSServer" `    //摄像机DNS服务器1
	CameraNTPServer    *string `json:"cameraNTPServer" `    //摄像机NTP服务器1
	StreamSource       *string `json:"streamSource" `       //视频流源 1
	CameraNVRChannelNo *string `json:"cameraNVRChannelNo" ` //通道号 1

	InOutFlag *int64 `json:"inOutFlag" ` //安装方向0-门外或进1-门内或出2-广场或其他开阔性的地方 1
}
type DeviceInfo struct {
	DeviceType  string  `json:"deviceType" validate:"required"` //1-网关 门禁 2-摄像机
	DeviceID    *string `json:"deviceID" `                      // 设备ID
	NewDeviceID *string `json:"newdeviceID" `                   // 新设备ID
}
type AccessInfo struct {
	AccessID        *string `json:"accessID"`         //ID
	SetNetworkWay   *int64  `json:"setNetworkWay" `   // 配网方式:0-手动配置 1-DCHP配置1
	Ip              *string `json:"ip" `              //ip 1
	SubnetMask      *string `json:"subnetMask"  `     // 子网掩码 1
	GateWay         *string `json:"gateWay" `         // 网关 1
	RtspUrl         *string `json:"rtspUrl" `         // 视频源流1
	NTPServer       *string `json:"NTPServer" `       // NTP服务器1
	NASServer       *string `json:"NASServer" `       // NAS服务器1
	IMGServer       *string `json:"IMGServer" `       //IMG服务器1
	OMServer        *string `json:"OMServer" `        // OMS服务器 1
	MQTTServer      *string `json:"MQTTServer" `      // MQTT服务器1
	DNSServer       *string `json:"DNSServer" `       // DNS服务器1
	BussinessServer *string `json:"BussinessServer" ` // 物管 1
}
type CameraInfo struct {
	CameraID         *string `json:"cameraID" `          //ID
	CameraIP         *string `json:"cameraIP" `          //ip1
	CameraSubnetMask *string `json:"cameraSubnetMask"  ` // 子网掩码 1
	FaceGateWay      *string `json:"faceGateWay" `       // 人脸网关1
	CameraOMServer   *string `json:"cameraOMServer" `    // OM服务器1
	CameraDNSServer  *string `json:"cameraDNSServer" `   //摄像机DNS服务器1
	CameraNTPServer  *string `json:"cameraNTPServer" `   //摄像机NTP服务器1
	StreamSource     *string `json:"streamSource" `      //视频流源 1
}
