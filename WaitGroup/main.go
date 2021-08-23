package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := new(sync.WaitGroup)
	for i := 0; i<5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()

	//variant with Mutex
	wg = new(sync.WaitGroup)
	mx := new(sync.Mutex)
	sum := 0
	MayWork(wg, mx, &sum)
	wg.Wait()
	fmt.Println(sum)

}

func MayWork(wg *sync.WaitGroup, mx *sync.Mutex, sum *int) {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mx.Lock()
			*sum +=1
			mx.Unlock()
		}()
	}
}

func work(){
	fmt.Println("Hello")
}