package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("I am learning Go")

	//? strings
	var nameOne string = "John"
	var nameTwo = "Doe"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Night"

	fmt.Println(nameFour)

	//? numbers

	var numOne int = 2
	var numTwo = 3
	numThree := 4
	fmt.Println(numOne, numTwo, numThree)

	//? boolean
	var isTrue bool //? default is false
	fmt.Println(isTrue) 
	isTrue = true
	fmt.Println(isTrue)
	isTrue = false
	fmt.Println(isTrue)

	//? arrays
	var arr []int = []int{1, 2, 3, 4, 5} //? this type can scale dynamically grow and shrink as an array
	fmt.Println(arr)
	var ages [4]int = [4]int{20, 25, 30} //? this type limit the size of the array and if the size is not enough it will add the value of ' 0 ' if there is needed for value in that array							
	fmt.Println(ages[0])
	fmt.Println(ages[1])
	fmt.Println(ages[2])
	ages[2] = 35
	fmt.Println(ages)
	fmt.Println(len(ages))
	fmt.Println(cap(ages))

	var names [4]string = [4]string{"John", "Doe", "Night", "Farmer"}
	fmt.Println(names)

	output := strings.Join(names[:], ", ")

	fmt.Println(output)

	greet(names[0])

	//? slices
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	slice = append(slice, 6)
	fmt.Println(slice)
	slice = append(slice, 7, 8, 9)
	fmt.Println(slice)
	slice = append(slice, 10)
	fmt.Println(slice)
	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println (slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = slice[:4]
	fmt.Println(slice)
	slice = slice[2:]
	fmt .Println(slice)

	//
	
}

func greet(name string){
	fmt.Println("Hello", name)
}