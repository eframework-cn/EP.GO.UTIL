//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 提供常用的IO操作，如读取文件、写入文件等功能.
package xfs

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/eframework-cn/EP.GO.UTIL/xlog"
	"github.com/eframework-cn/EP.GO.UTIL/xstring"
)

// 删除文件
//	path: 文件路径
func DeleteFile(path string) bool {
	if e := os.Remove(path); e != nil {
		xlog.Error("xfs.DeleteFile: %v", e)
		return false
	} else {
		return true
	}
}

// 判断文件是否存在
//	path: 文件路径
func FileExist(path string) bool {
	_, e := os.Stat(path)
	return e == nil
}

// 判断文件夹是否存在
//	path: 文件路径
//	createIsNotExist: 若不存在则创建
func PathExist(path string, createIsNotExist ...bool) bool {
	s, e := os.Lstat(path)
	if e != nil || !s.IsDir() {
		return false
	} else {
		if len(createIsNotExist) == 1 && createIsNotExist[0] {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return true
			}
		}
	}
	return false
}

// 根据文件名称获取该文件的文件夹名称
//	file: 文件路径
func Directory(file string) string {
	return filepath.Dir(file)
}

// 读取文件（以字符串形式返回）
//	file: 文件路径
func ReadString(file string) string {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		xlog.Error("xfs.ReadString: %v", e)
		return ""
	}
	return xstring.BytesToStr(b)
}

// 读取文件（以字节数组形式返回）
//	file: 文件路径
func ReadBytes(file string) []byte {
	b, e := ioutil.ReadFile(file)
	if e != nil {
		xlog.Error("xfs.ReadBytes: %v", e)
		return nil
	}
	return b
}

// 写入文件（字符串）
//	file: 文件路径
//	value: 文件内容
//	_mode: 文件模式
func WriteString(file string, value string, _mode ...os.FileMode) {
	mode := os.ModeAppend
	if len(_mode) == 1 {
		mode = _mode[0]
	}
	b := xstring.StrToBytes(value)
	ioutil.WriteFile(file, b, mode)
}

// 写入文件（字节数组）
//	file: 文件路径
//	value: 文件内容
//	_mode: 文件模式
func WriteBytes(file string, value []byte, _mode ...os.FileMode) {
	mode := os.ModeAppend
	if len(_mode) == 1 {
		mode = _mode[0]
	}
	ioutil.WriteFile(file, value, mode)
}
