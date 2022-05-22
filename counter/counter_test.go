package counter

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestCounter(t *testing.T) {
	counter := Init()

	counter.Set("a", 100)
	//log.Println(counter.Get("a"))
	assert.Equal(t, counter.Get("a"), int64(100))

	counter.Incr("a", 77)
	assert.Equal(t, counter.Get("a"), int64(177))

	counter.Init()
	assert.Equal(t, counter.Get("a"), int64(0))

}

func TestFlush2broker(t *testing.T) {
	counter := Init()
	counter.Flush2broker(2000, func() {
		log.Println("函数被调用了")
	})
}

//串行测试
func BenchmarkCounter(b *testing.B) {
	counter := Init()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		counter.Incr("a", 1)
		counter.Get("a")
	}
}

//并行测试
func BenchmarkCounterParallel(b *testing.B) {
	counter := Init()
	b.ResetTimer()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Incr("a", 1)
			counter.Get("a")
		}
	})
}
