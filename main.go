package main

import (
	"fmt"
	"sync"

	service "github.com/Dinesh789kumar12/CarBooking/service"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go service.CarLocation(&wg)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!!!")

}

// func main() {
// 	ticker := time.NewTicker(500 * time.Millisecond)
// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case t := <-ticker.C:
// 				fmt.Println("Tick at", t)
// 			}
// 		}
// 	}()

// 	time.Sleep(1600 * time.Millisecond)
// 	//ticker.Stop()
// 	done <- true
// 	fmt.Println("Ticker stopped")
// }
