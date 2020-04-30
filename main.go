package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

	sites := readLineFile()

	for i := 0; i < monitoring; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(sleepMonitoring * time.Second)
	}
}

func testSite(site string) {
	res, err := http.Get(site)

	if err != nil {
		fmt.Println("Error ->", err)
	} else {
		if res.StatusCode == 200 {
			fmt.Println(site, "-> Success", time.Now().Format(time.UnixDate))
		} else {
			fmt.Println(site, "->", res.StatusCode)
		}
	}
}

func readLineFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Error ->", err)
	} else {
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if err != nil && err != io.EOF {
				fmt.Println("Error ->", err)
			}

			sites = append(sites, line)

			if err == io.EOF {
				break
			}
		}
	}
	file.Close()

	return sites
}
