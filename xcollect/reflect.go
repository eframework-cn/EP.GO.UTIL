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

import (
	"reflect"

	"github.com/eframework-cn/EP.GO.UTIL/xlog"
)

// Deprecated: 判断集合是否存在元素
//	collection: 指针类型的map或slice
//	compare: 元素或对比函数
func Contains(collection interface{}, compare interface{}) bool {
	if collection == nil || compare == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	compareValue := reflect.ValueOf(compare)
	compareIsFunc := compareValue.Type().Kind() == reflect.Func
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Contains: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < collectionValue.Elem().Len(); i++ {
			obj := collectionValue.Elem().Index(i)
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					return true
				}
			} else {
				if ele == compare {
					return true
				}
			}
		}
	case reflect.Map:
		iter := collectionValue.Elem().MapRange()
		for iter.Next() {
			obj := iter.Value()
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					return true
				}
			} else {
				if ele == compare {
					return true
				}
			}
		}
	}
	return false
}

// Deprecated: 从集合中查找元素
//	collection: 指针类型的map或slice
//	compare: 元素或对比函数
func Find(collection interface{}, compare interface{}) interface{} {
	if collection == nil || compare == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	compareValue := reflect.ValueOf(compare)
	compareIsFunc := compareValue.Type().Kind() == reflect.Func
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Find: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < collectionValue.Elem().Len(); i++ {
			obj := collectionValue.Elem().Index(i)
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					return ele
				}
			} else {
				if ele == compare {
					return ele
				}
			}
		}
	case reflect.Map:
		iter := collectionValue.Elem().MapRange()
		for iter.Next() {
			obj := iter.Value()
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					return ele
				}
			} else {
				if ele == compare {
					return ele
				}
			}
		}
		return nil
	}
	return nil
}

// Deprecated: 从集合中索引元素
//	collection: 指针类型的map或slice
//	compare: 元素或对比函数
//	return: slice-number, map-key
func Index(collection interface{}, compare interface{}) interface{} {
	if collection == nil || compare == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	compareValue := reflect.ValueOf(compare)
	compareIsFunc := compareValue.Type().Kind() == reflect.Func
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Index: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < collectionValue.Elem().Len(); i++ {
			obj := collectionValue.Elem().Index(i)
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					return i
				}
			} else {
				if ele == compare {
					return i
				}
			}
		}
		return -1
	case reflect.Map:
		iter := collectionValue.Elem().MapRange()
		for iter.Next() {
			obj := iter.Value()
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					return iter.Key().Interface()
				}
			} else {
				if ele == compare {
					return iter.Key().Interface()
				}
			}
		}
	}
	return nil
}

// Deprecated: 在集合中新增元素
//	collection: 指针类型的slice
//	element: 元素
func Append(collection interface{}, element interface{}) bool {
	if collection == nil || element == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Append: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		array := []reflect.Value{reflect.ValueOf(element)}
		nslice := reflect.Append(collectionValue.Elem(), array...)
		collectionValue.Elem().Set(nslice)
		return true
	case reflect.Map:
	}
	return false
}

// Deprecated: 从集合中移除元素
//	collection: 指针类型的map或slice
//	compare: 元素或对比函数
func Remove(collection interface{}, compare interface{}) bool {
	if collection == nil || compare == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	compareValue := reflect.ValueOf(compare)
	compareIsFunc := compareValue.Type().Kind() == reflect.Func
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Remove: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		array := make([]reflect.Value, 0)
		sigRet := false
		for i := 0; i < collectionValue.Elem().Len(); i++ {
			obj := collectionValue.Elem().Index(i)
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					sigRet = true
				} else {
					array = append(array, obj)
				}
			} else {
				if ele == compare {
					sigRet = true
				} else {
					array = append(array, obj)
				}
			}
		}
		nslice := reflect.MakeSlice(collectionValue.Type().Elem(), 0, 0)
		nslice = reflect.Append(nslice, array...)
		collectionValue.Elem().Set(nslice)
		return sigRet
	case reflect.Map:
		iter := collectionValue.Elem().MapRange()
		for iter.Next() {
			obj := iter.Value()
			ele := obj.Interface()
			if compareIsFunc {
				params := []reflect.Value{obj}
				rets := compareValue.Call(params)
				if len(rets) == 1 &&
					rets[0].Kind() == reflect.Bool &&
					rets[0].Bool() {
					collectionValue.Elem().SetMapIndex(iter.Key(), reflect.Value{})
					return true
				}
			} else {
				if ele == compare {
					collectionValue.Elem().SetMapIndex(iter.Key(), reflect.Value{})
					return true
				}
			}
		}
	}
	return false
}

// Deprecated: 从集合中移除元素
//	collection: 支持指针类型的map或slice
//	compare: 索引或key
func Delete(collection interface{}, compare interface{}) bool {
	if collection == nil || compare == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	compareValue := reflect.ValueOf(compare)
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Delete: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		sigRet := false
		switch compare.(type) {
		case int, int32, int64, uint32, uint64,
			int16, uint16, int8, uint8:
			compareI := -1
			switch compare.(type) {
			case int:
				compareI = compare.(int)
				break
			case int32:
				compareI = int(compare.(int32))
				break
			case int64:
				compareI = int(compare.(int64))
				break
			case uint32:
				compareI = int(compare.(uint32))
				break
			case uint64:
				compareI = int(compare.(uint64))
				break
			case int16:
				compareI = int(compare.(int16))
				break
			case uint16:
				compareI = int(compare.(uint16))
				break
			case int8:
				compareI = int(compare.(int8))
				break
			case uint8:
				compareI = int(compare.(uint8))
				break
			}
			array := make([]reflect.Value, 0)
			for i := 0; i < collectionValue.Elem().Len(); i++ {
				obj := collectionValue.Elem().Index(i)
				if i == compareI {
					sigRet = true
				} else {
					array = append(array, obj)
				}
			}
			nslice := reflect.MakeSlice(collectionValue.Type().Elem(), 0, 0)
			nslice = reflect.Append(nslice, array...)
			collectionValue.Elem().Set(nslice)
			break
		default:
			xlog.Panic("xcollect.Delete: compare of slice must be interger")
			break
		}
		return sigRet
	case reflect.Map:
		value := collectionValue.Elem().MapIndex(compareValue)
		if value.IsValid() {
			collectionValue.Elem().SetMapIndex(compareValue, reflect.Value{})
			return true
		}
	}
	return false
}

// Deprecated: 清空集合中的元素
//	collection: 指针类型的map或slice
func Clear(collection interface{}) bool {
	if collection == nil {
		return false
	}
	collectionValue := reflect.ValueOf(collection)
	cKind := collectionValue.Kind()
	if cKind != reflect.Ptr {
		xlog.Panic("xcollect.Clear: collection must be pointer, this kind is", cKind)
		return false
	}
	switch collectionValue.Elem().Kind() {
	case reflect.Slice, reflect.Array:
		nslice := reflect.MakeSlice(collectionValue.Type().Elem(), 0, 0)
		collectionValue.Elem().Set(nslice)
		return true
	case reflect.Map:
		nmap := reflect.MakeMap(collectionValue.Type().Elem())
		collectionValue.Elem().Set(nmap)
		return true
	}
	return false
}
