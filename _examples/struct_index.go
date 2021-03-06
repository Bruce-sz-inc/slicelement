package main

import (
	"fmt"

	"github.com/1046102779/slicelement"
)

type Person struct {
	Name     string
	Age      int
	Children []string
}

func main() {
	data := []*Person{
		&Person{
			Name:     "John",
			Age:      29,
			Children: []string{"David", "Lily", "Bruce Lee"},
		},
		&Person{
			Name:     "Joe",
			Age:      18,
			Children: []string{},
		},
	}
	elem := 18
	index, err := slicelement.GetIndex(data, elem, "Age")
	if err != nil {
		//fmt.Println(errors.Cause(err).Error())
		fmt.Println(err.Error())
	}
	fmt.Println("index=", index)
	// output: index=1
}
