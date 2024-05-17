package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func day1() {
	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	fmt.Println("Hello, World!")

	scanner := bufio.NewScanner(file)
	ans := 0
	numMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for scanner.Scan() {
		line := scanner.Text()

		var first, second int
		foundFirst, foundSecond := false, false
		var first_str string
		var second_str string

		for _, char := range line {
			first_str += string(char)
			if len(first_str) >= 3 {
				temp_first3 := first_str[len(first_str)-3:]
				value, exists := numMap[temp_first3]

				if exists {
					first = value
					foundFirst = true
					break
				}
			}
			if len(first_str) >= 4 {
				temp_first4 := first_str[len(first_str)-4:]
				value, exists := numMap[temp_first4]

				if exists {
					first = value
					foundFirst = true
					break
				}
			}
			if len(first_str) >= 5 {
				temp_first5 := first_str[len(first_str)-5:]
				value, exists := numMap[temp_first5]

				if exists {
					first = value
					foundFirst = true
					break
				}
			}
			if unicode.IsDigit(char) {
				first = int(char - '0')
				foundFirst = true
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			char := rune(line[i])
			second_str = string(char) + second_str
			if len(second_str) >= 3 {
				temp_first3 := second_str[:3]
				value, exists := numMap[temp_first3]

				if exists {
					second = value
					foundSecond = true
					break
				}
			}
			if len(second_str) >= 4 {
				temp_first4 := second_str[:4]
				value, exists := numMap[temp_first4]

				if exists {
					second = value
					foundSecond = true
					break
				}
			}
			if len(second_str) >= 5 {
				temp_first5 := second_str[:5]
				value, exists := numMap[temp_first5]

				if exists {
					second = value
					foundSecond = true
					break
				}
			}
			if unicode.IsDigit(char) {
				second = int(char - '0')
				foundSecond = true
				break
			}
		}

		if foundFirst && foundSecond {
			ret := first*10 + second
			fmt.Println(ret)
			ans += ret
		} else {
			fmt.Println("Could not find two digits in the line.")
		}
	}

	fmt.Println("Final answer:", ans)
}

func main() {
	day1()
}
