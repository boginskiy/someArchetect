package maps

import (
	"sync"
	"sync/atomic"
)

type CountMaper interface {
	SnapShot() map[string]*int64
	Put(key string, value int64)
	Take(key string) int64
}

type CounterMap struct {
	sync.RWMutex
	store map[string]*int64
}

func NewCounterMap() *CounterMap {
	return &CounterMap{
		store: make(map[string]*int64, 10),
	}
}

func (c *CounterMap) SnapShot() map[string]*int64 {
	deepCopyMap := make(map[string]*int64)

	c.Lock()
	defer c.Unlock()

	for k, v := range c.store {
		newV := new(int64)
		*newV = *v

		deepCopyMap[k] = newV
	}

	clear(c.store)
	return deepCopyMap
}

func (c *CounterMap) Put(key string, value int64) {
	c.Lock()

	cntP, ok := c.store[key]
	if !ok {
		newCnt := new(int64)
		c.store[key] = newCnt
		cntP = newCnt
	}

	c.Unlock()
	atomic.AddInt64(cntP, value)
}

func (c *CounterMap) Take(key string) int64 {
	c.RLock()

	value, ok := c.store[key]
	if !ok || value == nil {
		return 0
	}

	return atomic.LoadInt64(value)
}
