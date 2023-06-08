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

// 从数组/切片中索引元素（[]int64）
//	arr: 数组/切片
//	ele: 元素
func IndexForInt64(arr []int64, ele int64) int {
	if arr != nil {
		for k, v := range arr {
			if v == ele {
				return k
			}
		}
	}
	return -1
}

// 数组/切片是否存在元素（[]int64）
//	arr: 数组/切片
//	ele: 元素
func ContainsForInt64(arr []int64, ele int64) bool {
	if arr != nil {
		return IndexForInt64(arr, ele) >= 0
	}
	return false
}

// 从数组/切片中移除元素（[]int64）
//	arr: 数组/切片
//	ele: 元素
func RemoveForInt64(arr []int64, ele int64) []int64 {
	if arr != nil {
		for {
			idx := IndexForInt64(arr, ele)
			if idx >= 0 {
				arr = append(arr[:idx], arr[idx+1:]...)
			} else {
				break
			}
		}
	}
	return arr
}

// 从数组/切片中移除元素（[]int64）
//	arr: 数组/切片
//	index: 索引
func DeleteForInt64(arr []int64, ele int) []int64 {
	if arr != nil {
		if ele < len(arr) {
			arr = append(arr[:ele], arr[ele+1:]...)
		}
	}
	return arr
}

// 在数组/切片中新增元素（[]int64）
//	arr: 数组/切片
//	ele: 元素
func AppendForInt64(arr []int64, ele int64) []int64 {
	return append(arr, ele)
}
