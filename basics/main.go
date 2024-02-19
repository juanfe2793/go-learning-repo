package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*
	 Pointer logic
	*/

	name := "Felipe"  // Implicit data type. The := is used only the first time the variale is declared.
	fmt.Println(name) // print felipe

	pointer := &name //POinter is targeting to the value in the memory. If the value of the variable "name" changes, pointer will tracked it.

	fmt.Println(pointer, *pointer) // Print the MEM address, the value stored in MEM (felipe)

	name = "juan"
	fmt.Println(pointer, *pointer) // Print the MEM address, the value stored in MEM (juan)

	var lastName string = "Gomez" // It is possible to be more verbose with the declaration of the variable.
	fmt.Println(lastName)

	age := 30 // There are multiple int options: int8, int16, int32, int64. Use the one that better fits your needs.
	age = age + 4

	fmt.Println(age)
	fmt.Println(age - 15)

	fmt.Println(name+" "+lastName, age) // Concatenation of strings

	fmt.Println(reflect.TypeOf(age)) // Print the data type of the variable

	var myFloat float32 = 3.14 // There are multiple float options: float32, float64. Use the one that better fits your needs.
	fmt.Println(myFloat)

	fmt.Println(float64(age) + float64(myFloat)) // It is necessary to convert the data type to be able to sum the variables.

	/*
		Constants
	*/

	const pi float64 = 3.1416 // Constants are declared with the keyword const and don't requires to be used.

	/*
		Arrays
	*/

	var myArray [5]int // Arrays are fixed size. The size is defined in the declaration.
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	myArray[3] = 4
	myArray[4] = 5

	fmt.Println(myArray)

	/*
		Maps
	*/

	myMap := make(map[string]int) // Maps are dynamic size. The size is not defined in the declaration.
	myMap[name] = age
	myMap[lastName] = age - 10
	myMap["Holi"] = 100

	fmt.Println(myMap)

	/*
		Loops
	*/

	for index := 0; index < len(myArray); index++ {
		fmt.Println(myArray[index])
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println(funcName())

	/*
		Structures for data types
	*/

	type Person struct { // Requires to start with a capital letter to be able to use it in other packages.
		name string
		age  int
	}

	person := Person{name: name, age: age}
	fmt.Println(person)
}

func funcName() string {

	return "Felipe"

}
