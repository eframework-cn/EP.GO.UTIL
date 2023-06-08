//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 提供对json序列化/反序列化的支持.
package xjson

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/eframework-cn/EP.GO.UTIL/xstring"
)

var (
	TO_OBJ_INVALID_PARAM  = errors.New("xjson.ToObj invalid param")      // 错误：不合法的参数
	TO_OBJ_PARSE_ERROR    = errors.New("xjson.ToObj parse content fail") // 错误：解析内容失败
	TO_OBJ_UNSUPPORT_DATA = errors.New("xjson.ToObj unsupport type")     // 错误：不支持的类型
)

func encodeNoEscape(v interface{}) ([]byte, error) {
	switch v.(type) {
	case string:
		return xstring.StrToBytes(v.(string)), nil
	}
	buff := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buff)
	encoder.SetEscapeHTML(false)
	if e := encoder.Encode(v); e != nil {
		return nil, e
	}
	bytes := buff.Bytes()
	size := len(bytes)
	if bytes[size-1] == 10 { // 去除换行符
		bytes = append(bytes[:size-1])
	}
	return bytes, nil
}

// 对象转字符串输出
func ToPrint(v interface{}) string {
	str, _ := ToString(v)
	return str
}

// 对象转字符串
func ToString(v interface{}) (string, error) {
	b, e := encodeNoEscape(v)
	if e != nil {
		return "", fmt.Errorf("xjson.ToString: %v", e)
	}
	return xstring.BytesToStr(b), nil
}

// 对象转字节数组
func ToByte(v interface{}) ([]byte, error) {
	b, e := encodeNoEscape(v)
	if e != nil {
		return nil, fmt.Errorf("xjson.ToByte: %v", e)
	}
	return b, nil
}

// 字符串/字节数组转对象（指针）
func ToObj(data interface{}, obj interface{}) error {
	if data == nil || obj == nil {
		return TO_OBJ_INVALID_PARAM
	}
	switch data.(type) {
	case string:
		str := data.(string)
		if str == "null" {
			return nil
		}
		if e := json.Unmarshal(xstring.StrToBytes(str), obj); e != nil {
			return e
		}
		return nil
	case []byte:
		str := xstring.BytesToStr(data.([]byte))
		if str == "null" {
			return nil
		}
		if e := json.Unmarshal(data.([]byte), obj); e != nil {
			return e
		}
		return nil
	}
	return TO_OBJ_UNSUPPORT_DATA
}
