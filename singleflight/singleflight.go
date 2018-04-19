package singleflight

import (
	"sync"
)

type call struct {
	value interface{}
	err error
	wg sync.WaitGroup
}

type Group struct {
	lock sync.Mutex
	calls map[string]*call
}

func (g *Group)Do(key string, fn func()(val interface{}, err error)) (interface{}, error){
	g.lock.Lock()
	if g.calls == nil {
		g.calls = make(map[string]*call)
	}
	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait()
		return c.value, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.calls[key] = c
	g.lock.Unlock()

	c.value, c.err = fn()
	c.wg.Done()

	g.lock.Lock()
	delete(g.calls, key)
	g.lock.Unlock()
	return c.value, c.err
}

