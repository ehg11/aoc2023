package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
    rhs := strings.Split(line, ":")[1]
    rhs = strings.TrimSpace(rhs)

    strValues := strings.Split(rhs, " ")
    var values []int
    for _, str := range(strValues) {
        str = strings.TrimSpace(str)
        num, err := strconv.Atoi(str)
        if err != nil {
            continue
        }
        values = append(values, num)
    }

    return values
}

func parseLine2(line string) int {
    rhs := strings.Split(line, ":")[1]
    rhs = strings.TrimSpace(rhs)

    strValues := strings.Split(rhs, " ")
    numStr := ""
    for _, str := range(strValues) {
        str = strings.TrimSpace(str)
        numStr += str
    }
    num, _ := strconv.Atoi(numStr)

    return num
}

func countWinningMethods(time int, distance int) int {
    winningMethods := 0
    for i := 0; i < time; i++ {
        movingTime := time - i
        if movingTime * i > distance {
            winningMethods += 1
        }
    }

    return winningMethods
}

func main() {
    filepath := "d6"
    file, _ := os.Open(filepath)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var buffer []string

    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }
    fmt.Println(buffer)

    times := parseLine2(buffer[0])
    distances := parseLine2(buffer[1])

    fmt.Println(countWinningMethods(times, distances))
}
