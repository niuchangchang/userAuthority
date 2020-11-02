/*
 * @Description: 备注
 * @Author: Dawn
 * @Date: 2019-08-06 14:57:46
 * @LastEditTime: 2019-08-07 13:10:05
 * @LastEditors: Dawn
 */
package models

import (
	"github.com/wangcong0918/sunrise/sql_orm"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
)

type ShortEngine struct {
	sql_orm.ShortEngine
}

func GetModelsEngine() (err error) {
	engineConn := sql_orm.EngineCon
	Engine, err = engineConn.GetOrmEngine()

	if err != nil {
		return err
	}

	return nil
}

func (s ShortEngine) GetShortModelsEngine() (engine *xorm.Engine, err error) {
	engine, err = s.GetShortEngine()
	if err != nil {
		return nil, err
	}

	return engine, nil
}
