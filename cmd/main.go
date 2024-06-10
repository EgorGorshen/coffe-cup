package main

import (
	particles "coffee-cup/pkg"
	"fmt"
	"time"
)

func main() {
	coffee := particles.NewCoffee(63, 25, 9.0)
	timer := time.NewTicker(100 * time.Millisecond)
    coffee.Start()
	for {
        <-timer.C
        fmt.Print("\033[H\033[2J")
		coffee.Update()
		fmt.Println(coffee.Display())
	}
}
