package cache

import (
	"sync"
)

var Cache *cache

type cache struct {
	progress bool
	mutex    sync.Mutex
	task     chan struct{}
	process  func()
}

func init() {
	Cache = &cache{
		task: make(chan struct{}, 1),
	}
	Cache.Monitor()
}

func (c *cache) AddEvent() {
	c.mutex.Lock()
	if c.progress == true {
		c.mutex.Unlock()
		return
	}
	c.progress = true
	c.mutex.Unlock()
	c.task <- struct{}{}
}

func (c *cache) SetProcess(process func()) {
	c.process = process
}

func (c *cache) Monitor() {
	go c.Execute()
}

func (c *cache) Execute() {
	for range c.task {
		c.mutex.Lock()
		c.progress = false
		c.mutex.Unlock()
		c.process()
	}
}
