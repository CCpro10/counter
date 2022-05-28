package counter

import (
	"sync"
	"time"
)

//作业功能：实现一个计数器模块，不依赖外部三方模块和存储，要求进程内协程安全、异步、高性能按指标
//Key - value 统计。使用示例：
//counter .Init()
//counter . Incr (" get . called ",123) counter . Incr (" get . called ",456)

var Counter *counter

type counter struct {
	m sync.Map
}

type Node struct {
	Key   string
	Count int64
}

func Init() *counter {
	return &counter{
		sync.Map{},
	}
}

func (c *counter) Init() {
	c.m = sync.Map{}
}

func (c *counter) Get(key string) (count int64) {

	value, ok := c.m.Load(key)
	if !ok {
		return 0
	}
	return value.(int64)
}

func (c *counter) GetAll() []Node {
	var values []Node
	c.m.Range(func(key, value interface{}) bool {
		values = append(values, Node{
			Key:   key.(string),
			Count: value.(int64),
		})
		return true
	})
	return values
}

func (c *counter) Set(key string, count int64) {
	c.m.Store(key, count)
}

func (c *counter) Delete(key string) {
	c.m.Delete(key)
}

//Return the Count after Incr
//If the Key is not exist ,Incr will set Key to num
func (c *counter) Incr(key string, num int64) (count int64) {
	value, ok := c.m.Load(key)
	if !ok {
		c.m.Store(key, num)
		return num
	}
	c.m.Store(key, value.(int64)+num)
	return value.(int64) + num
}

//counter .Flush2broker(5000, FuncCbFlush )
// cal FlushCb every 5s and reset counter
func (c *counter) Flush2broker(ms int64, FuncCbFlush func()) {
	duration := time.Millisecond * time.Duration(ms)
	ticker := time.NewTicker(duration)
	for range ticker.C {
		FuncCbFlush()
		c.Init()
	}
}
