package singleton

import (
	"fmt"
	"sync"
)

var once sync.Once

type AlternativeSingleton struct {
	callCount int
}

var alternativeInstance *AlternativeSingleton

func getAlternativeInstance() *AlternativeSingleton {
	if alternativeInstance == nil {
		once.Do(
			func() {
				alternativeInstance = &AlternativeSingleton{callCount: 1}
				fmt.Printf("Creating alternative instance. Call count: %v\n", alternativeInstance.callCount)
			})
	} else {
		alternativeInstance.callCount++
		fmt.Printf("Alternative already exists. Call count: %v\n", alternativeInstance.callCount)
	}

	return alternativeInstance
}
