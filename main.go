package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoring = 3
const sleepMonitoring = 5

func main() {
	showIntro()

	// infinity loop
	for {
		showMenu()

		switch readCommand() {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Show logs...")
		case 0:
			fmt.Println("Exit, Bye!")
			os.Exit(0)
		default:
			fmt.Println("I don't understand your command typed!")
			os.Exit(-1)
		}
	}
}

func showIntro() {
	version := 1.0
	fmt.Println("web monitor version", version)
}

func showMenu() {
	fmt.Println("\n[1] Start monitor")
	fmt.Println("[2] Show logs")
	fmt.Println("[0] Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := []string{"http://github.com", "http://gitlab.com", "http://bitbucket.com"}

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(sleepMonitoring * time.Second)
	}
}

func testSite(site string) {
	res, _ := http.Get(site)
	if res.StatusCode == 200 {
		fmt.Println(site, "-> Success", time.Now().Format(time.UnixDate))
	} else {
		fmt.Println(site, "->", res.StatusCode)
	}
}
