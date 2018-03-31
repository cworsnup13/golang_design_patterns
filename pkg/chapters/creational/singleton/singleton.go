package main

import (
	"sync"
	"fmt"
)

type singleton struct {
}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func GetAndPrintAddrDemo(wg *sync.WaitGroup) {
	s := GetInstance()
	fmt.Printf("%p\n", s)
	wg.Done()
	return
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go GetAndPrintAddrDemo(&wg)
	}
	wg.Wait()
}
