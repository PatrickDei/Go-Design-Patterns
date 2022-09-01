package singleton

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

type Singleton struct {
	callCount int
}

var instance *Singleton

func getInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{callCount: 1}
			fmt.Printf("Creating the instance. Call count: %v\n", instance.callCount)
		} else {
			instance.callCount++
			fmt.Printf("Already exists. Call count: %v\n", instance.callCount)
		}
	} else {
		instance.callCount++
		fmt.Printf("Already exists. Call count: %v\n", instance.callCount)
	}

	return instance
}
