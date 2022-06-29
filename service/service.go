package service

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Car = []string{"Audi", "Maruti", "Suzuki", "Honda",
	"Benz"}

func CarLocation(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		for _, id := range Car {
			a, b := RandomNumberGenerator()
			fmt.Printf("Car: %v \t Location is {%v %v}\n", id, a, b)
		}
		time.Sleep(30 * time.Second)
		CarBooked("Audi") // Car Booking(delete from slice)
	}
}
func RandomNumberGenerator() (x int, y int) {
	rand.Seed(time.Now().UnixNano())

	a := rand.Intn(100) + 1
	b := rand.Intn(100) + 1

	return a, b
}

func removeCar(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func CarBooked(s string) {
	Car = removeCar(Car, s)
}

func RealTimeData() {

}
