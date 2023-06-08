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
	"syscall"
	"time"
	"unsafe"

	"github.com/eframework-cn/EP.GO.UTIL/xstring"
	"github.com/eframework-cn/EP.GO.UTIL/xtime"

	"github.com/lxn/win"
	"golang.org/x/sys/windows"
)

var (
	libkernel32      *windows.LazyDLL
	setConsoleTitleW *windows.LazyProc
)

func init() {
	libkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	setConsoleTitleW = libkernel32.NewProc("SetConsoleTitleW")
}

func setConsoleTitleWFunc(name *uint16) {
	syscall.Syscall(setConsoleTitleW.Addr(), 1, uintptr(unsafe.Pointer(name)), 0, 0)
}

func SetCmdTitle(title string) {
	if s, e := syscall.UTF16PtrFromString(title); e == nil {
		setConsoleTitleWFunc(s)
	}
}

func DisableClose() {
	title := xstring.Format("ESVR-%v", xtime.GetMicrosecond())
	if s, e := syscall.UTF16PtrFromString(title); e == nil {
		setConsoleTitleWFunc(s)
		time.Sleep(time.Millisecond * 500)
		hwnd := win.FindWindow(nil, s)
		menu := win.GetSystemMenu(hwnd, false)
		win.RemoveMenu(menu, 0xF060, 0x0)
	}
}
