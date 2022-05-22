package counter

import (
	"sync"
	"time"
)

//作业功能：实现一个计数器模块，不依赖外部三方模块和存储，要求进程内协程安全、异步、高性能按指标
//key - value 统计。使用示例：
//counter .Init()
//counter . Incr (" get . called ",123) counter . Incr (" get . called ",456)

type counter struct {
	m     map[string]int64
	mutex sync.RWMutex
}

func Init() *counter {
	return &counter{
		make(map[string]int64),
		sync.RWMutex{},
	}
}

func (c *counter) Init() {
	c.m = make(map[string]int64)
}

func (c *counter) Get(key string) int64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.m[key]
}

func (c *counter) Set(key string, count int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.m[key] = count
}

func (c *counter) Incr(key string, num int64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.m[key] += num
}

//counter .Flush2broker(5000, FuncCbFlush )
// cal FlushCb every 5s and reset counter
func (c *counter) Flush2broker(ms int, FuncCbFlush func()) {
	duration := time.Millisecond * time.Duration(ms)
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.Init()
		FuncCbFlush()
	}
}
