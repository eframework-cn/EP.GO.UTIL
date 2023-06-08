//-----------------------------------------------------------------------//
//                     GNU GENERAL PUBLIC LICENSE                        //
//                        Version 2, June 1991                           //
//                                                                       //
// Copyright (C) EFramework, https://eframework.cn, All rights reserved. //
// Everyone is permitted to copy and distribute verbatim copies          //
// of this license document, but changing it is not allowed.             //
//                   SEE LICENSE.md FOR MORE DETAILS.                    //
//-----------------------------------------------------------------------//

// 时间常用函数封装，如获取时间戳、时间格式化等.
package xtime

import (
	"time"
)

const (
	STRING_TIME_FORMAT_FULL = "2006-01-02 15:04:05 +0800 CST" // 时间格式化模板（标准）
	STRING_TIME_FORMAT_LITE = "2006-01-02 15:04:05"           // 时间格式化模板（简易）
	STRING_TIME_FORMAT_FILE = "2006-01-02_15_04_05"           // 时间格式化模板（文件）
)

const (
	SEC_1   int = 1       // 1秒
	SEC_5   int = 5       // 5秒
	SEC_10  int = 10      // 10秒
	SEC_15  int = 15      // 15秒
	SEC_20  int = 20      // 20秒
	SEC_25  int = 25      // 25秒
	SEC_30  int = 30      // 30秒
	SEC_35  int = 35      // 35秒
	SEC_40  int = 40      // 40秒
	SEC_45  int = 45      // 45秒
	SEC_50  int = 50      // 50秒
	SEC_55  int = 55      // 55秒
	MIN_1   int = 60      // 1分钟
	MIN_2   int = 120     // 2分钟
	MIN_3   int = 180     // 3分钟
	MIN_4   int = 240     // 4分钟
	MIN_5   int = 300     // 5分钟
	MIN_6   int = 360     // 6分钟
	MIN_7   int = 420     // 7分钟
	MIN_8   int = 480     // 8分钟
	MIN_9   int = 540     // 9分钟
	MIN_10  int = 600     // 10分钟
	MIN_12  int = 720     // 12分钟
	MIN_15  int = 900     // 15分钟
	MIN_20  int = 1200    // 20分钟
	MIN_25  int = 1500    // 25分钟
	MIN_30  int = 1800    // 30分钟
	MIN_35  int = 2100    // 35分钟
	MIN_40  int = 2400    // 40分钟
	MIN_45  int = 2700    // 45分钟
	MIN_50  int = 3000    // 50分钟
	MIN_55  int = 3300    // 55分钟
	HOUR_1  int = 3600    // 1小时
	HOUR_2  int = 7200    // 2小时
	HOUR_3  int = 10800   // 3小时
	HOUR_4  int = 14400   // 4小时
	HOUR_5  int = 18000   // 5小时
	HOUR_6  int = 21600   // 6小时
	HOUR_7  int = 25200   // 7小时
	HOUR_8  int = 28800   // 8小时
	HOUR_9  int = 32400   // 9小时
	HOUR_10 int = 36000   // 10小时
	HOUR_11 int = 39600   // 11小时
	HOUR_12 int = 43200   // 12小时
	HOUR_13 int = 46800   // 13小时
	HOUR_14 int = 50400   // 14小时
	HOUR_15 int = 54000   // 15小时
	HOUR_16 int = 57600   // 16小时
	HOUR_17 int = 61200   // 17小时
	HOUR_18 int = 64800   // 18小时
	HOUR_19 int = 68400   // 19小时
	HOUR_20 int = 72000   // 20小时
	HOUR_21 int = 75600   // 21小时
	HOUR_22 int = 79200   // 22小时
	HOUR_23 int = 82800   // 23小时
	DAY_1   int = 86400   // 1天
	DAY_2   int = 172800  // 2天
	DAY_3   int = 259200  // 3天
	DAY_4   int = 345600  // 4天
	DAY_5   int = 432000  // 5天
	DAY_6   int = 518400  // 6天
	DAY_7   int = 604800  // 7天
	DAY_8   int = 691200  // 8天
	DAY_9   int = 777600  // 9天
	DAY_10  int = 864000  // 10天
	DAY_15  int = 1296000 // 15天
	DAY_20  int = 1728000 // 20天
	DAY_30  int = 2592000 // 30天
)

// 获取时间戳（微秒）
func GetMicrosecond() int {
	ltime := time.Now().UnixNano() / 1e3
	return int(ltime)
}

// 获取时间戳（毫秒）
func GetMillisecond() int {
	ltime := time.Now().UnixNano() / 1e6
	return int(ltime)
}

// 获取时间戳（秒）
func GetTimestamp() int {
	ltime := time.Now().Unix()
	time.Unix(0, 0).Format("")
	return int(ltime)
}

// 获取当前时间
func NowTime() time.Time {
	return time.Now()
}

// 时间戳（秒）转时间（object）
//	timestamp: 时间戳
func ToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

// 获取距离零点的时间（秒）
//	timestamp: 指定时间，若未指定，则使用当前时间
func TimeToZero(timestamp ...int) int {
	t := 0
	if len(timestamp) == 1 {
		t = timestamp[0]
	} else {
		t = GetTimestamp()
	}
	return DAY_1 - (t+HOUR_8)%DAY_1
}

// 获取零点的时间戳（秒）
//	timestamp: 指定时间，若未指定，则使用当前时间
func ZeroTime(timestamp ...int) int {
	t := 0
	if len(timestamp) == 1 {
		t = timestamp[0]
	} else {
		t = GetTimestamp()
	}
	return t - (t+HOUR_8)%DAY_1
}

// 时间格式化
//	timestamp: 时间戳
//	format: 格式化模板
func Format(timestamp int, format string) string {
	return time.Unix(int64(timestamp), 0).Format(format)
}
