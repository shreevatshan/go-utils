package maps

import (
	"sync"
	"time"
)

type ExpiringItem struct {
	Value any
	timer *time.Timer
}

type ExpiringMap struct {
	sync.Map
	mu sync.Mutex
}

func NewExpiringMap() *ExpiringMap {
	return &ExpiringMap{}
}

func (em *ExpiringMap) StoreWithExpiry(key any, value any, duration time.Duration) {
	em.mu.Lock()
	defer em.mu.Unlock()

	if t, ok := em.Load(key); ok {
		if _, ok := t.(*ExpiringItem); ok {
			t.(*ExpiringItem).timer.Stop()
		}
	}

	if duration == 0 {
		em.Store(key, value)
		return
	}

	timer := time.AfterFunc(duration, func() {
		em.Delete(key)
	})

	em.Store(key, &ExpiringItem{value, timer})
}

func (em *ExpiringMap) Fetch(key any) (value any, ok bool) {
	v, ok := em.Load(key)
	if !ok {
		return nil, false
	}

	if _, ok := v.(*ExpiringItem); !ok {
		return v, true
	}

	return v.(*ExpiringItem).Value, true
}
