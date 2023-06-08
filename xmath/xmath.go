//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 封装了一些常用的数学函数.
package xmath

import (
	"encoding/binary"
	"math/rand"
)

const (
	INT8_MIN   = -0x7f - 1
	INT16_MIN  = -0x7fff - 1
	INT32_MIN  = -0x7fffffff - 1
	INT64_MIN  = -0x7fffffffffffffff - 1
	INT8_MAX   = 0x7f
	INT16_MAX  = 0x7fff
	INT32_MAX  = 0x7fffffff
	INT64_MAX  = 0x7fffffffffffffff
	UINT8_MAX  = 0xff
	UINT16_MAX = 0xffff
	UINT32_MAX = 0xffffffff
	UINT64_MAX = 0xffffffffffffffff
)

var (
	BYTE_ORDER binary.ByteOrder = &binary.LittleEndian // 字节序（小端）
)

// 最大值
func MaxValue(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 最小值
func MinValue(a int, b int) int {
	if a >= b {
		return b
	}
	return a
}

// 随机数
//	min: 左区间
//	max: 右区间
func RandInt(min int, max int) int {
	if min >= max {
		return max
	}
	return rand.Intn(max-min) + min
}

// uint32转字节数组
func Uint32ToBytes(value uint32) []byte {
	bytes := make([]byte, 4)
	BYTE_ORDER.PutUint32(bytes, value)
	return bytes
}

// 字节数组转uint32
func Uint32FromBytes(value []byte) uint32 {
	return BYTE_ORDER.Uint32(value)
}
