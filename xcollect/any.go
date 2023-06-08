//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 封装了any/int/int32/int64/string数组/切片的常用操作.
package xcollect

// 从数组/切片中索引元素（[]interface{}）
//	arr: 数组/切片
//	ele: 元素
func IndexForAny(arr []interface{}, ele interface{}) int {
	if arr != nil {
		switch ele.(type) {
		case func(ele interface{}) bool:
			for k, v := range arr {
				if ele.(func(ele interface{}) bool)(v) {
					return k
				}
			}
			break
		case interface{}:
			for k, v := range arr {
				if v == ele {
					return k
				}
			}
			break
		}
	}
	return -1
}

// 数组/切片是否存在元素（[]interface{}）
//	arr: 数组/切片
//	ele: 元素
func ContainsForAny(arr []interface{}, ele interface{}) bool {
	if arr != nil {
		return IndexForAny(arr, ele) >= 0
	}
	return false
}

// 从数组/切片中移除元素（[]interface{}）
//	arr: 数组/切片
//	ele: 元素
func RemoveForAny(arr []interface{}, ele interface{}) []interface{} {
	if arr != nil {
		switch ele.(type) {
		case func(ele interface{}) bool:
			for i := 0; i < len(arr); {
				v := arr[i]
				if ele.(func(ele interface{}) bool)(v) {
					arr = append(arr[:i], arr[i+1:]...)
				} else {
					i++
				}
			}
			break
		case interface{}:
			for {
				idx := IndexForAny(arr, ele)
				if idx >= 0 {
					arr = append(arr[:idx], arr[idx+1:]...)
				} else {
					break
				}
			}
			break
		}
	}
	return arr
}

// 从数组/切片中移除元素（[]interface{}）
//	arr: 数组/切片
//	index: 索引
func DeleteForAny(arr []interface{}, index int) []interface{} {
	if arr != nil {
		if index < len(arr) {
			arr = append(arr[:index], arr[index+1:]...)
		}
	}
	return arr
}

// 在数组/切片中新增元素（[]interface{}）
//	arr: 数组/切片
//	ele: 元素
func AppendForAny(arr []interface{}, ele interface{}) []interface{} {
	return append(arr, ele)
}
