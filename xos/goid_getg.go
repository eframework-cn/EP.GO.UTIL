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
	"fmt"
	"unsafe"
)

const (
	GOID_OFFSET uintptr = 152
)

func init() {
	fmt.Println("GOID_OFFSET: ", GOID_OFFSET)
}

func goid_getg() int64 {
	g := getg()
	if g == 0 {
		return 0
	} else {
		p := (*int64)(unsafe.Pointer(g + GOID_OFFSET))
		return *p
	}
}

func getg() uintptr
