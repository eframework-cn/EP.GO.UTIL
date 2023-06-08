# xlog (derive from beego/logs[v1.12.0])
* 移除了logs/alils和logs/es
* 修改了logMsg(struct)，支持格式化输出，新增toString函数
* 修改了BeeLogger(struct)，新增了closed相关逻辑