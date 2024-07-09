package main

import (
	"fmt"
	"os"
	"time"
)

func Select(get chan int, quit chan struct{}) {

	// Select block is like switch case only
	for {
		time.Sleep(time.Second)
		select {
		case <-get:
			fmt.Println("Got something")
		case <-quit:
			fmt.Println("Quit signal recived Quiting....")
			os.Exit(0)
		}
	}

}

func main() {
	get := make(chan int)
	quit := make(chan struct{})

	go func() {
		get <- 1
		quit <- struct{}{} // Empty struct
	}()

	go Select(get, quit)
	select {} // Waits indefinitely
}
