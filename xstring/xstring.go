//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 字符串常用函数封装，如转数字、截取、修建、替换、切割等.
package xstring

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// 字符串转int
func ToInt(str string) int {
	itr, _ := strconv.Atoi(str)
	return itr
}

// int转字符串
func ToString(itr int) string {
	str := strconv.Itoa(itr)
	return str
}

// int64转字符串
func Int64ToString(itr int64) string {
	str := strconv.FormatInt(itr, 10)
	return str
}

// 保留指定位数的小数（参数为float32或float64）（默认保留两位小数）
//	fixed: 指定小数点位数
func ToFixed(float interface{}, fixed ...int) string {
	tfixed := 2
	if len(fixed) == 1 {
		tfixed = fixed[0]
	}
	if tfixed < 0 {
		tfixed = 0
	}
	if float != nil {
		switch float.(type) {
		case float32, float64:
			return fmt.Sprintf("%."+ToString(tfixed)+"f", float)
		}
	}
	return ""
}

// 字符串分割
//	sep: 分割符
func Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

// 找到指定字符的索引
func IndexOf(str string, of string) int {
	return strings.Index(str, of)
}

// 找到指定字符的索引（后）
func LastIndexOf(str string, of string) int {
	return strings.LastIndex(str, of)
}

// 是否以指定字符起始
func StartWith(str string, of string) bool {
	return strings.HasPrefix(str, of)
}

// 是否以指定字符结束
func EndWith(str string, of string) bool {
	return strings.HasSuffix(str, of)
}

// 是否包含指定字符
func Contains(str string, of string) bool {
	return strings.Contains(str, of)
}

// 是否为空
func IsEmpty(str string) bool {
	return str == ""
}

// 截取
func Sub(str string, from int, to int) string {
	rs := []rune(str)
	length := len(rs)
	if from < 0 || to < 0 || from > to {
		return ""
	}
	if to > length {
		to = length
	}
	return string(rs[from:to])
}

// 替换所有指定字符
func Replace(str string, from string, to string) string {
	return strings.ReplaceAll(str, from, to)
}

// 剔除多余的空格
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// 字符串转字节数组
func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// 字节数组转字符串
func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 字符串格式化
func Format(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
