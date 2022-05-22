package main

import (
	"log"
	"main/counter"
)

func main() {
	c := counter.Init()

	c.Set("key", 100)
	log.Println(c.Get("key"))
	c.Incr("key", 77)
	log.Println(c.Get("key"))
	c.Init()
	go c.Flush2broker(3000, func() {
		log.Println("函数被调用了一次")
	})
	select {}
}
