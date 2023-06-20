//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 提供了日志打印（分层&异步）、日志保存（分块&异步）等功能.
package xlog

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	MAX_CHAN_LEN = 300000 // 通道缓存数量
)

// 日志配置
type LogCfg struct {
	Filename string `json:"filename,omitempty"` // 日志名称
	MaxLines int    `json:"maxlines,omitempty"` // 最大行数
	MaxSize  int    `json:"maxsize,omitempty"`  // 最大字节数
	Daily    bool   `json:"daily,omitempty"`    // 跨日自动保存
	MaxDays  int64  `json:"maxdays,omitempty"`  // 最多支持天数
	Rotate   bool   `json:"rotate,omitempty"`
	Level    int    `json:"level,omitempty"` // 日志层级
	Perm     string `json:"perm,omitempty"`  // 文件权限
}

// 初始化
//	cfgs: 配置
func Init(cfgs map[string]*LogCfg) *BeeLogger {
	setLogger := func(name string, cfg *LogCfg) {
		r, _ := json.Marshal(cfg)
		e := beeLogger.SetLogger(name, string(r))
		if e != nil {
			fmt.Printf("xlog.Init: set logger failed, error=%v", e)
		}
	}

	level := LevelEmergency
	for k, v := range cfgs {
		setLogger(k, v)
		if v.Level > level {
			level = v.Level
		}
	}

	beeLogger.SetLevel(level)

	Async(MAX_CHAN_LEN)

	return beeLogger
}

// 清空缓冲区
func Flush() {
	Emergency("xlog.Flush: GLogger has been flushed")
	beeLogger.Flush()
}

// 关闭日志服务
func Close() {
	Emergency("xlog.Close: GLogger has been closed")
	beeLogger.Close()
}

// 缓冲区数量
func ChanSize() int {
	return len(beeLogger.msgChan)
}

// 触发异常（Crash）
func Panic(v ...interface{}) {
	s := fmt.Sprintf(strings.Repeat("%v ", len(v)), v...)
	panic(s)
}
