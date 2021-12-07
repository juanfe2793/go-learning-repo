package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	// Command line parameters definition. The variable is setup as a parameter if you use flag package.

	path := flag.String("path", "system.log", "This is the path of the log file you want to verify")
	level := flag.String("level", "com.apple.asl", "The log level you want to filter")
	flag.Parse() //This will populate the parameters

	file, err := os.Open(*path)
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
		if strings.Contains(str, *level) {
			fmt.Println(str)
		}
	}

}
