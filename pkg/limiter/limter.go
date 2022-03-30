package limiter

import (
	"golang.org/x/time/rate"
	"log"
	"sync"
	"time"
)

type Limiters struct {
	Limiters map[string]*Limiter
	Lock     sync.Mutex
}

type Limiter struct {
	limiter    *rate.Limiter
	lastAccess time.Time
	//key        string
}

func NewLimitersCollection() *Limiters {
	return &Limiters{
		Limiters: make(map[string]*Limiter),
		Lock:     sync.Mutex{},
	}
}

func (l *Limiter) Allow() bool {
	access := time.Now()
	l.lastAccess = access
	return l.limiter.Allow()
}

func (ls *Limiters) GetLimiter(r rate.Limit, b int, key string) *Limiter {
	ls.Lock.Lock()
	defer ls.Lock.Unlock()
	if limiter, ok := ls.Limiters[key]; ok {
		return limiter
	}

	newLimiter := &Limiter{
		limiter:    rate.NewLimiter(r, b),
		lastAccess: time.Now(),
		//key:        key,
	}

	ls.Limiters[key] = newLimiter
	return newLimiter
}

func (ls *Limiters) ClearNotUseLimiter(sec time.Duration) {
	for {
		time.Sleep(sec) //for now just set 1-minutes for testing
		//for all limiter
		for key, l := range ls.Limiters {
			if time.Now().Sub(l.lastAccess) > sec {
				ls.Lock.Lock()
				delete(ls.Limiters, key)
				log.Printf("limiter for ip:%v is removed", key)
				ls.Lock.Unlock()
			}
		}
	}
}
