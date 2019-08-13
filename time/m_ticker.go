package time

import (
	"github.com/astaxie/beego"
	"runtime"
	"time"
)

// 定时器函数
func timeTicker(f func(), t time.Duration) {
	// 将线程绑定到系统线程
	runtime.LockOSThread()
	ticker := time.NewTicker(t)
	defer func() {
		ticker.Stop()
		runtime.UnlockOSThread()
		// panic处理，打印错误信息
		if err := recover(); err != nil {
			beego.Error("定时器故障! 捕获到了panic错误：", err)
		}
	}()

	// 遍历
	for _ = range ticker.C {
		f()
	}
}
