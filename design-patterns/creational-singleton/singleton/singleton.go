package singleton

import (
	"sync"
	"time"
)

type Singleton struct {
	Time string
}

var (
	once     sync.Once
	instance Singleton
)

func New() *Singleton {
	once.Do(func() {
		instance = Singleton{
			Time: time.Now().String(),
		}
	})

	return &instance
}
