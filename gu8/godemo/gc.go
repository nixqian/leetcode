package godemo

import (
	"runtime"
	"runtime/debug"
)

func test() {
	runtime.GC() // 手动触发gc
	debug.SetGCPercent(100)
	// GOGC 环境变量 100 增长100%触发, 12 增长12%触发GC
	// 按时间设置: 默认2分钟
}
