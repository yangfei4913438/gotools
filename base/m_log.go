package base

import "github.com/astaxie/beego/logs"

//WEB日志
func SetLog() *logs.BeeLogger {
	//日志模块初始化
	log := logs.NewLogger()

	//日志设置命令行输出
	log.SetLogger("console")

	//增加日志的文件输出
	log.SetLogger(logs.AdapterFile, `{"filename":"logs/http_project.log","level":7,"daily":true,"maxdays":7}`)

	//显示日志行号以及错误信息所在文件名称
	log.EnableFuncCallDepth(true)

	//日志设置完成,返回设置好的日志模块
	return log
}
