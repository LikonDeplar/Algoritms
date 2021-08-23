package main

import "fmt"

func main() {

	//first variant
	done := make( chan struct{})
	go func(){
		defer close(done)
		work()
	}()
	<-done

	//second variant
	<-MayWork()


}

func work() {
	fmt.Println("Hello")
}

func MayWork() <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		fmt.Println("Hello")

	}()
	return done
}