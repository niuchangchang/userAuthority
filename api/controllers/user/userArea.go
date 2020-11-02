/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-19 14:13:01
 * @LastEditTime: 2019-09-27 09:51:26
 * @LastEditors: Dawn
 */
package user

import (
	"reflect"
)

// 获取系统用户对应区域权限
func GetErectorVillageList(userID string) () {

}
func Convert(vinfo interface{}) map[string]interface{} {
	t := reflect.TypeOf(vinfo)
	v := reflect.ValueOf(vinfo)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
