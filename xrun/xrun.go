//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 协程上层封装，包括异常捕捉、闪退重启、消耗时间、调用堆栈等功能.
package xrun

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/eframework-cn/EP.GO.UTIL/xfs"
	"github.com/eframework-cn/EP.GO.UTIL/xlog"
	"github.com/eframework-cn/EP.GO.UTIL/xtime"
)

const (
	UNKONWN_SOURCE = "[?]"
)

// 获取函数调用堆栈信息
//	stack: 堆栈层级
//	fullpath: 全路径
func Caller(stack int, fullpath bool) string {
	if pc, file, line, ok := runtime.Caller(stack + 1); ok {
		if fullpath {
			return fmt.Sprintf("[%s:%d (0x%v)]", file, line, pc)
		} else {
			return fmt.Sprintf("[%s:%d (0x%v)]", runtime.FuncForPC(pc).Name(), line, pc)
		}
	}
	return UNKONWN_SOURCE
}

// 获取错误堆栈信息
//	stack: 堆栈层级
//	err: 错误信息
func StackTrace(stack int, err interface{}) (string, int) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	start := stack + 1
	count := stack
	fmt.Fprintf(buf, "    skip %v\n", stack)
	for i := start; ; i++ {
		line := Caller(i, true)
		if line == UNKONWN_SOURCE {
			break
		}
		count++
		fmt.Fprintf(buf, "    %v\n", line)
	}
	return buf.String(), count
}

// 计算函数执行消耗的时间，在起始处调用defer Elapse(stack)()
//	stack: 堆栈层级（0为当前层）
func Elapse(stack int, callback ...func()) func() {
	start := xtime.GetMillisecond()
	return func() {
		end := xtime.GetMillisecond()
		elapse := end - start
		if stack < 0 {
			stack = 0
		}
		caller := Caller(stack+1, false)
		xlog.Notice("xrun.Elapse%v: start-%v, end-%v, elapsed-%vms", caller, start, end, elapse)
		if len(callback) == 1 {
			callback[0]()
		}
	}
}

func doCaught(exit bool, stack int, handler ...func(string, int)) {
	if err := recover(); err != nil {
		str, count := StackTrace(stack+1, err)
		fname := fmt.Sprintf("log/crash/%v.log", xtime.Format(xtime.GetTimestamp(), xtime.STRING_TIME_FORMAT_FILE))
		xfs.PathExist(xfs.Directory(fname), true)
		xfs.WriteString(fname, str)
		xlog.Critical(str)
		if len(handler) == 1 {
			handler[0](str, count)
		}
		if exit {
			xlog.Critical("exit now!")
			xlog.Close()
			os.Exit(1)
		}
	}
}

// 捕捉goroutine的异常
//	exit: 是否结束进程
//	handler: 回调
func Caught(exit bool, handler ...func(string, int)) {
	doCaught(exit, 0, handler...)
}

// 协程封装器（panic不会引起crash，recover后该goroutine结束）
func Exec(fun interface{}, params ...interface{}) {
	defer Caught(false)
	vf := reflect.ValueOf(fun)
	vps := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		param := params[i]
		vps[i] = reflect.ValueOf(param)
	}
	vf.Call(vps)
}

func doRun(fun interface{}, stack int, params ...interface{}) {
	defer doCaught(false, stack, func(s string, i int) {
		doRun(fun, i, params...)
	})
	vf := reflect.ValueOf(fun)
	vps := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		vps[i] = reflect.ValueOf(params[i])
	}
	vf.Call(vps)
}

// 协程封装器（panic不会引起crash，recover后重启该goroutine）
func Run(fun interface{}, params ...interface{}) {
	doRun(fun, 0, params...)
}
