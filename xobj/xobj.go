//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 面向对象编程.
package xobj

// 对象基类
type OBJECT struct {
	CHILD interface{} `orm:"-" json:"-"` // 实例指针对象
}

// 构造函数
//	CHILD: 实例指针对象
func (this *OBJECT) CTOR(CHILD interface{}) {
	this.CHILD = CHILD
}
