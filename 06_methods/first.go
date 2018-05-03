package main

import "fmt"

type Computer struct {
	Brand string
	Model string
	Price int
}

// TODO: Create a "Describe" method that prints all the properties
func (c Computer) Describe() string {
	//return fmt.Println()
	return fmt.Sprintf("Brand:%v/Model:%v/Price:%v", c.Brand, c.Model, c.Price)
}

func main() {
	computer := Computer{
		Brand: "Apple",
		Model: "Macbook",
		Price: 1000,
	}

	fmt.Println(computer.Describe())

}
