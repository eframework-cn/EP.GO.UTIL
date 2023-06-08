//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

package xcollect

// 从数组/切片中索引元素（[]int）
//	arr: 数组/切片
//	ele: 元素
func IndexForInt(arr []int, ele int) int {
	if arr != nil {
		for k, v := range arr {
			if v == ele {
				return k
			}
		}
	}
	return -1
}

// 数组/切片是否存在元素（[]int）
//	arr: 数组/切片
//	ele: 元素
func ContainsForInt(arr []int, ele int) bool {
	if arr != nil {
		return IndexForInt(arr, ele) >= 0
	}
	return false
}

// 从数组/切片中移除元素（[]int）
//	arr: 数组/切片
//	ele: 元素
func RemoveForInt(arr []int, ele int) []int {
	if arr != nil {
		for {
			idx := IndexForInt(arr, ele)
			if idx >= 0 {
				arr = append(arr[:idx], arr[idx+1:]...)
			} else {
				break
			}
		}
	}
	return arr
}

// 从数组/切片中移除元素（[]int）
//	arr: 数组/切片
//	index: 索引
func DeleteForInt(arr []int, index int) []int {
	if arr != nil {
		if index < len(arr) {
			arr = append(arr[:index], arr[index+1:]...)
		}
	}
	return arr
}

// 在数组/切片中新增元素（[]int）
//	arr: 数组/切片
//	ele: 元素
func AppendForInt(arr []int, ele int) []int {
	return append(arr, ele)
}
