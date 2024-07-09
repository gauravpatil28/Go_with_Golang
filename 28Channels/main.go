package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GO lang")

	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	/* Channels in go lang these are the way through which the go routines communicate to each other we can make
	   a chanel simply by using the make statement and can be of any type int,string,..  and we assume the chanel
	   to be a box so when we write like this chan<- we are inserting values in the chanel and when we write like
	   this <-chan we are taking values out of it. and yes chanel only works when there is another go routine listning
	   to it else it will give and error naming DEADLOCK! . and yes how many times you want it to print values you should
	   give it command to print the values too. else again the deadlock will appear

	*/
	wg.Add(2)

	//Recieve only Channel
	go func(ch chan<- int, wg *sync.WaitGroup) {

		val, IsChanelOpen := <-myCh
		fmt.Println(IsChanelOpen) //GIves you bool value that is the channel open or not
		fmt.Println(val)
		wg.Done()

	}(myCh, wg)

	// Output only channel
	go func(ch <-chan int, wg *sync.WaitGroup) {

		myCh <- 5
		//myCh <- 8
		close(myCh) //This closes the channel

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
