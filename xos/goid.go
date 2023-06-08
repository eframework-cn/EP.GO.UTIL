//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 获取运行时元信息（协程ID），提供了若干对系统API的封装调用，如设置应用标题等.
package xos

// 获取当前的GoroutineID
//	notice: 应当尽早使用Context替换GOID逻辑
// 共有三个版本获取GOID（性能比为: 90:3:1）
//	1.stack: 通过堆栈信息获取
//	2.go-tls: go 1.17.*版本无法使用go-tls，参考https://github.com/huandu/go-tls，[20211125待修复]
//	3.getg: 通过汇编代码获取g结构体，再确定goid的offset，用offset获取goid，参考https://github.com/v2pro/plz
func GoID(bystack ...bool) int64 {
	if len(bystack) == 0 || bystack[0] == false {
		return goid_getg()
	} else {
		return goid_stack()
	}
}
