package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 3
const sleepMonitoring = 5
const fileNameLog = "sites.log"
const fileSiteList = "sites.txt"

func main() {
	showIntro()

	// infinity loop
	for {
		showMenu()

		switch readCommand() {
		case 1:
			startMonitoring()
		case 2:
			printLogs()
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
	fmt.Println("")

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
			fmt.Println(site, "-> Success", time.Now().Format("2006/01/02 03:04:05 PM"))
		} else {
			fmt.Println(site, "->", res.StatusCode)
		}
		saveLog(site, res.StatusCode)
	}
}

func readLineFile() []string {
	var sites []string
	file, err := os.Open(fileSiteList)

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

func saveLog(site string, status int) {
	file, err := os.OpenFile(fileNameLog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Error ->", err)
	} else {
		file.WriteString(time.Now().Format("2006/01/02 03:04:05 PM") + " | " + site + "-> Status:" + strconv.Itoa(status) + "\n")
	}

	file.Close()
}

func printLogs() {
	fmt.Println("Print logs...")
	fmt.Println("")

	file, err := ioutil.ReadFile(fileNameLog)

	if err != nil {
		fmt.Println("Error ->", err)
	} else {
		fmt.Println(string(file))
	}
}
