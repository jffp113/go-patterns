package singleton

import "sync"

var (
	once     sync.Once
	instance *singleton
)

type singleton struct {
	a int
}

func GetInstance() *singleton {

	once.Do(func() {
		instance = &singleton{a: 10}
	})

	return instance
}

func GetInstanceWithIf() *singleton {

	if instance == nil {
		once.Do(func() {
			instance = &singleton{a: 10}
		})
	}

	return instance
}

func resetSingleton() {
	once = sync.Once{}
	instance = nil
}
