package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
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

        var first, second int
        foundFirst, foundSecond := false, false

        for _, char := range line {
            if unicode.IsDigit(char) {
                first = int(char - '0')
                foundFirst = true
                break
            }
        }

        for i := len(line) - 1; i >= 0; i-- {
            char := rune(line[i])
            if unicode.IsDigit(char) {
                second = int(char - '0')
                foundSecond = true
                break
            }
        }

        if foundFirst && foundSecond {
            ret := first*10 + second
            fmt.Printf("First digit: %d, Second digit: %d, Result: %d\n", first, second, ret)
            ans += ret
        } else {
            fmt.Println("Could not find two digits in the line.")
        }
    }

    fmt.Println("Final answer:", ans)
}
