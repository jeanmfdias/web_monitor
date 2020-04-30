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
}
