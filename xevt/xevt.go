//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 提供事件的注册、注销、通知等功能.
package xevt

import (
	"github.com/eframework-cn/EP.GO.UTIL/xcollect"
	"github.com/eframework-cn/EP.GO.UTIL/xlog"
	"github.com/eframework-cn/EP.GO.UTIL/xrun"
)

// 消息句柄
type IHandler interface {
	Handle(reply *EvtReply, receiver interface{}, param interface{})
}

// 消息中心
type EvtMgr struct {
	HID   int                // 自增ID
	Mutil bool               // 是否支持多个收听者
	Evts  map[int]*EvtEntity // 事件映射
}

// 消息实体
type EvtEntity struct {
	ID   int          // 事件ID
	Hnds []*HndEntity // 句柄列表
}

// 句柄实体
type HndEntity struct {
	ID   int      // 句柄ID
	Func IHandler // 句柄函数
}

// 消息响应
type EvtReply struct {
	Result interface{} // 透传参数
	Pend   chan int    // 阻塞标识
}

// 新建消息响应
//	pend: 阻塞标识
func NewEvtReply(pend ...chan int) *EvtReply {
	this := new(EvtReply)
	if len(pend) == 1 {
		this.Pend = pend[0]
	}
	return this
}

// 完成消息响应
//	result: 透传参数
func (this *EvtReply) Done(result ...interface{}) {
	if len(result) == 1 {
		this.Result = result[0]
	}
	if this.Pend != nil {
		close(this.Pend)
	}
}

// 新建消息中心
//	multi: 是否支持多个收听者
func NewEvtMgr(mutli bool) *EvtMgr {
	return &EvtMgr{0, mutli, make(map[int]*EvtEntity)}
}

// 清除消息注册
func (this *EvtMgr) Clear() {
	this.Evts = make(map[int]*EvtEntity)
}

// 获取消息实体
//	id: 消息ID
func (this *EvtMgr) Get(id int) *EvtEntity {
	return this.Evts[id]
}

// 注册消息
//	id: 消息ID
//	handler: 回调函数
func (this *EvtMgr) Reg(id int, handler IHandler) int {
	if nil == handler {
		xlog.Error("EvtMgr.Reg: nil handler, id=%v", id)
		return -1
	}
	entity := this.Evts[id]
	if entity == nil {
		entity = new(EvtEntity)
		entity.ID = id
		entity.Hnds = make([]*HndEntity, 0)
		this.Evts[id] = entity
	}
	if this.Mutil == false && len(entity.Hnds) > 1 {
		xlog.Error("EvtMgr.Reg: not support multi-register, id=%v", id)
		return -1
	}
	this.HID++
	hnd := new(HndEntity)
	hnd.ID = this.HID
	hnd.Func = handler
	entity.Hnds = append(entity.Hnds, hnd)
	return hnd.ID
}

// 注销消息
//	id: 消息ID
//	hid: 句柄ID
func (this *EvtMgr) Unreg(id int, hid int) bool {
	entity := this.Evts[id]
	if entity == nil {
		return false
	} else {
		return xcollect.Remove(&entity.Hnds, func(ele interface{}) bool { return ele.(*HndEntity).ID == hid })
	}
}

// 通知消息
//	id: 消息ID
//	reply: 响应对象
//	param1: 参数1
//	param2: 参数2
func (this *EvtMgr) Notify(id int, reply *EvtReply, param1 interface{}, param2 interface{}) bool {
	defer xrun.Caught(false)
	entity := this.Get(id)
	if entity == nil {
		if id < 0 {
			// reserve msg id
		} else {
			xlog.Warn("xevt.Notify: no handler of msgid=%v found.", id)
		}
		return false
	}
	for _, h := range entity.Hnds {
		if h != nil && h.Func != nil {
			h.Func.Handle(reply, param1, param2)
		}
	}
	return true
}
