/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-05 10:28:04
 * @LastEditTime: 2019-08-19 15:28:08
 * @LastEditors: Dawn
 */
package middleware

import (
	"net/http"
	"github.com/wangcong0918/sunrise"
)

// Response make middleware 返回的resbody
func ResponseResult(t *sunrise.Context, resBody interface{}) {
	t.JSON(http.StatusOK, resBody)
}
