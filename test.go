package main

import (
    "fmt"
    "os"
	"log"
	"bufio"
	"unicode"
)

func main() {
	file, err := os.Open("a.txt")
	if err != nil {
        log.Fatalf("failed to open file: %s", err)
    }
    defer file.Close()
    fmt.Println("Hello, World!")

	scanner := bufio.NewScanner(file)
	ans := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		sum := 0
		for _, char := range line {
			if unicode.IsDigit(char) {
				sum = sum * 10 + int(char - '0')
			}
		}
		fmt.Println(sum)
		ans += sum

	}

	fmt.Println(ans)

}
