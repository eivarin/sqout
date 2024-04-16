package TimersMap

import (
	"sync"
	"time"
)

type TimersMap struct {
	timers map[string]*time.Timer
	lock  sync.RWMutex
}

func NewTimersMap() TimersMap {
	return TimersMap{
		timers: make(map[string]*time.Timer),
		lock: sync.RWMutex{},
	}
}

func (t *TimersMap) WaitFor(key string, seconds int) {
	t.lock.Lock()
	delete(t.timers, key)
	t.timers[key] = time.NewTimer(time.Duration(seconds) * time.Second)
	t.lock.Unlock()
	t.lock.RLock()
	<-t.timers[key].C
	t.lock.RUnlock()
}