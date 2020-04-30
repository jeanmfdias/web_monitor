package main

import "fmt"

func main() {
	var command int

	version := 1.0

	fmt.Println("web monitor version", version)

	fmt.Println("[1] Start monitor")
	fmt.Println("[2] Show logs")
	fmt.Println("[0] Exit")

	fmt.Scan(&command)

	switch command {
	case 1:
		fmt.Println("Monitoring...")
	case 2:
		fmt.Println("Show logs:")
	case 0:
		fmt.Println("Exit, Bye!")
	default:
		fmt.Println("I don't understand your command typed!")
	}
}
