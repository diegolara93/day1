package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("[0-9]")
var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var pattern = strings.Join(words, "|")

func addCount(text []string) int {
	if len(text) > 1 {
		for i := range text {
			arr := re.FindAllString(text[i], -1)
			twoNumbers := arr[0] + arr[len(arr)-1]
			fmt.Println(twoNumbers)
			convInt, err := strconv.Atoi(twoNumbers)
			if err != nil {
				panic(err)
			}
			return convInt + addCount(text[i+1:])
		}
	}
	return 0
}

func main() {
	content, error := os.ReadFile("d:/go/advent-of-code/day1/static/input.txt")
	contentString := string(content)
	splitString := strings.Split(contentString, "\n")
	fmt.Println((addCount(splitString)))
	if error != nil {
		fmt.Println("Error", error)
		os.Exit(1)
	}
}
