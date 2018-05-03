package main

import (
	"fmt"
	"time"
)

type Computer struct {
	Brand string
	Model string
	Price int
}

func (c *Computer) Describe() string {
	return fmt.Sprintf("Brand:%v/Model:%v/Price:%v", c.Brand, c.Model, c.Price)
}

// TODO: Make a StartTimer function that...
//  - Sleeps for a given duration
//  - Prints a message when done
func StartTimer(s string) {
	t := 3 * time.Second
	for i := 0; i < 5; i++ {
		time.Sleep(t)
		fmt.Println(s)
	}
}

func main() {
	computer := Computer{
		Brand: "Apple",
		Model: "Macbook",
		Price: 1000,
	}
	StartTimer(computer.Describe())
	//computer.StartTimer(t)
}
