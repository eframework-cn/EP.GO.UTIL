//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 提供解析json、xml、ini等格式的配置文件解析.
package xconfig

import (
	"encoding/json"
	"encoding/xml"

	"github.com/eframework-cn/EP.GO.UTIL/xfs"
	"github.com/eframework-cn/EP.GO.UTIL/xlog"
)

type parseFunc func(data []byte, v interface{}) error

func doParse(parse parseFunc, file string, out interface{}) error {
	xlog.Notice("xconfig.doParse: load file at %v", file)

	data := xfs.ReadBytes(file)
	if data != nil {
		if err := parse(data, out); err != nil {
			xlog.Error("xconfig.doParse: parse failed, err=%v", file, err)
			return err
		}
	}

	return nil
}

// 反序列化Json
//	file: json文件
//	out: obj对象
func JsonToObj(file string, obj interface{}) error {
	return doParse(json.Unmarshal, file, obj)
}

// 反序列化Xml
//	file: xml文件
//	out: obj对象
func XmlToObj(file string, obj interface{}) error {
	if err := doParse(xml.Unmarshal, file, obj); err != nil {
		return err
	}
	return nil
}
