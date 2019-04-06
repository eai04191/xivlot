package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/Songmu/prompter"
)

func main() {
	username := os.Getenv("USERNAME")
	path := prompter.Prompt("Enter ACT log path", "C:\\Users\\"+username+"\\AppData\\Roaming\\Advanced Combat Tracker\\FFXIVLogs\\")
	name := prompter.Prompt("Character Name", "Your Name")

	var rolls []int

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("Path Open Error")
	}

	for _, f := range files {
		log.Print("Log Path: " + path + f.Name())
		f, err := os.Open(path + f.Name())
		if err != nil {
			log.Fatal("Log Open Error")
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			// 一行ずつ処理

			r := regexp.MustCompile(name + `は.*に(GREED|NEED)のダイスで(\d{1,3})を出した。.*$`)
			if r.MatchString(scanner.Text()) {
				match := r.FindStringSubmatch(scanner.Text())
				log.Print("Found: " + match[2])
				dice, _ := strconv.Atoi(match[2])
				rolls = append(rolls, dice)
			}
		}

	}

	sum := 0
	for _, x := range rolls {
		sum += x
	}

	average := sum / len(rolls)

	fmt.Println("")
	fmt.Println("Result")
	fmt.Println("Rolls Count:", len(rolls), "times")
	fmt.Println("Average:", average)
	prompter.Prompt("", "")
}
