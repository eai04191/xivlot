package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"

	"github.com/Songmu/prompter"
	"github.com/dustin/go-humanize"
)

func sum(s []int) int {
	var sum int
	for _, x := range s {
		sum += x
	}
	return sum
}

func max(s []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	return s[0]
}

func min(s []int) int {
	sort.Sort(sort.IntSlice(s))
	return s[0]
}

func main() {
	scan()
}

func scan() {
	username := os.Getenv("USERNAME")
	lang := prompter.Choose("FFXIV Language", []string{"JA", "EN"}, "JA")
	path := prompter.Prompt("ACT Log Path", "C:\\Users\\"+username+"\\AppData\\Roaming\\Advanced Combat Tracker\\FFXIVLogs\\")
	var name string
	if lang != "EN" {
		name = prompter.Prompt("Character Name", "Your Name")
	}

	regexpPatterns := map[string]string{
		"JA": name + `は.*に(GREED|NEED)のダイスで(\d{1,3})を出した。.*$`,
		"EN": `You roll (Greed|Need) on the .*\. (\d{1,3})!.*$`,
	}

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
			r := regexp.MustCompile(regexpPatterns[lang])
			if r.MatchString(scanner.Text()) {
				match := r.FindStringSubmatch(scanner.Text())
				log.Print("Found: " + match[2])
				dice, _ := strconv.Atoi(match[2])
				rolls = append(rolls, dice)
			}
		}

	}

	show(rolls)
}

func show(rolls []int) {
	fmt.Println("")
	fmt.Println("Result")

	fmt.Printf("%20s", "Total Rolls: ")
	fmt.Println(humanize.Comma(int64(len(rolls))), "times")

	fmt.Printf("%20s", "Sum: ")
	fmt.Println(humanize.Comma(int64(sum(rolls))))

	fmt.Printf("%20s", "Max: ")
	fmt.Println(max(rolls))

	fmt.Printf("%20s", "Min: ")
	fmt.Println(min(rolls))

	fmt.Printf("%20s", "Average: ")
	fmt.Printf("%.2f\n", float64(sum(rolls))/float64(len(rolls)))

	fmt.Println("")
	prompter.Prompt("Enter to Exit.", "")
}
