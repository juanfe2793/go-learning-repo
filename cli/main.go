package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("system.log")
	if err != nil {
		log.Fatal(err)
	}
	// This will close the file after finish the execution
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(str, "com.apple.asl") {
			fmt.Println(str)
		}
	}

}
