package main

import (
	"fmt"
	"sync"
	"time"
)

var done = false

func read(name string, c *sync.Cond) {
	fmt.Println(name, "starts reading")

	c.L.Lock()
	for !done {
		c.Wait() // 等待发出通知
	}
	fmt.Println(name, "readed")
	c.L.Unlock()
}

func write(name string, c *sync.Cond) {
	fmt.Println(name, "starts writing")
	time.Sleep(100 * time.Millisecond)

	c.L.Lock()
	done = true // 设置条件变量
	c.L.Unlock()

	fmt.Println(name, "wakes all")
	c.Broadcast() // 通知所有观察者
}

func main() {
	cond := sync.NewCond(&sync.Mutex{}) // 创建时传入一个互斥锁

	// 3 个观察者
	go read("reader1", cond)
	go read("reader3", cond)
	go read("reader2", cond)

	time.Sleep(time.Second) // 模拟延时

	write("writer-1", cond) // 发出通知

	time.Sleep(time.Second) // 模拟延时
}
