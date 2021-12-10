package main

import "fmt"

func main() {

	/*
	*Pointer logic
	 */

	name := "Felipe"  // Implicit data type. The := is used only the first time the variale is declared.
	fmt.Println(name) // print felipe

	pointer := &name //POinter is targeting to the value in the memory. If the value of the variable "name" changes, pointer will tracked it.

	fmt.Println(pointer, *pointer) // Print the MEM address, the value stored in MEM (felipe)

	name = "juan"
	fmt.Println(pointer, *pointer) // Print the MEM address, the value stored in MEM (juan)

}
