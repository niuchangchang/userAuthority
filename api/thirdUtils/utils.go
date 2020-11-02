/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-06 10:31:18
 * @LastEditTime: 2019-09-10 14:50:55
 * @LastEditors: Dawn
 */
package thirdUtils

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	uuid4, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return ""
	}
	uuids := uuid4.String()
	//通过函数进行替换
	re3, _ := regexp.Compile("-")
	//把匹配的所有字符a替换成b
	rep2 := re3.ReplaceAllString(uuids, "")
	// fmt.Println(rep2)
	return rep2
}

func DecHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

//第一个传数据库查出来的[]map[string] interface{} 第二个传入需要绑定的interface
func Convert(Map interface{}, pointer interface{}) {
	// reflect.Ptr类型 *main.Person
	pointertype := reflect.TypeOf(pointer)
	// reflect.Value类型
	pointervalue := reflect.ValueOf(pointer)
	// struct类型
	structType := pointertype.Elem()
	// 将interface{}类型的map转换为  map[string]interface{}
	m := Map.(map[string]interface{})
	// 遍历结构体字段
	for i := 0; i < structType.NumField(); i++ {
		// 获取指定字段的反射值
		f := pointervalue.Elem().Field(i)
		//获取struct的指定字段
		stf := structType.Field(i)
		// 获取tag
		name := strings.Split(stf.Tag.Get("json"), ",")[0]
		// 判断是否为忽略字段
		if name == "-" {
			continue
		}
		// 判断是否为空，若为空则使用字段本身的名称获取value值
		if name == "" {
			name = stf.Name
		}
		//获取value值
		v, ok := m[name]
		if !ok {
			continue
		}

		//获取指定字段的类型
		kind := pointervalue.Elem().Field(i).Kind()
		// 若字段为指针类型
		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = f.Type().Elem().Kind()
		}
		// 设置对应字段的值
		switch kind {
		case reflect.Int:
			res, _ := strconv.ParseInt(fmt.Sprint(v), 10, 64)
			pointervalue.Elem().Field(i).SetInt(res)
		case reflect.String:
			pointervalue.Elem().Field(i).SetString(fmt.Sprint(v))
		}
	}
}

type Data struct {
	FirstClass  string `json:"first_class"`  // 一级
	SecondClass string `json:"second_class"` // 二级
	ThirdClass  string `json:"third_class"`  //三级
}

