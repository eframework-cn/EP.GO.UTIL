//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

package xos

import (
	"runtime"
)

var (
	pc_lookup = make(map[uintptr]int8, 17)
)

const (
	bitWidth       = 4
	stackBatchSize = 16
)

func get_stack(offset, amount int) (stack []uintptr, next_offset int) {
	stack = make([]uintptr, amount)
	stack = stack[:runtime.Callers(offset, stack)]
	if len(stack) < amount {
		return stack, 0
	}
	return stack, offset + len(stack)
}

func read_stack() (tag uint, ok bool) {
	var current_tag uint
	offset := 0
	for {
		batch, next_offset := get_stack(offset, stackBatchSize)
		for _, pc := range batch {
			val, ok := pc_lookup[pc]
			if !ok {
				continue
			}
			if val < 0 {
				return current_tag, true
			}
			current_tag <<= bitWidth
			current_tag += uint(val)
		}
		if next_offset == 0 {
			break
		}
		offset = next_offset
	}
	return 0, false
}

func goid_stack() int64 {
	i, _ := read_stack()
	return int64(i)
}
