package server

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"sync"
	"time"
)

type SentCache struct {
	group []*openwechat.SentMessage
	time  time.Time
}

type ConcurrentMap struct {
	mu    *sync.RWMutex
	cache map[string]SentCache
}

var Buckets = ConcurrentMap{
	mu:    new(sync.RWMutex),
	cache: make(map[string]SentCache),
}

//每分钟监听一次缓存,防止内存过大

func (m *ConcurrentMap) StartExpirationCheck(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("test", m)
			m.CheckAndDeleteExpired()
		}
	}
}

func (m *ConcurrentMap) CheckAndDeleteExpired() {
	m.mu.Lock()
	defer m.mu.Unlock()
	now := time.Now()
	for key, val := range m.cache {
		if now.Sub(val.time) >= (2 * time.Minute) {
			delete(m.cache, key)
		}
	}
}
