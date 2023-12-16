package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkSubstrNum(substr string) int {
    stringNums := map[string]int{
        "one":      1,
        "two":      2,
        "three":    3,
        "four":     4,
        "five":     5,
        "six":      6,
        "seven":    7,
        "eight":    8,
        "nine":     9,
    }

    foundNum := -1
    for stringNum, num := range stringNums {
        if strings.HasPrefix(substr, stringNum) {
            foundNum = num
            break
        }
    }

    return foundNum
}

func findNums(line string) (int, int) {
    leftNum := -1
    rightNum := -1
    for i := 0; i < len(line); i++ {
        newNum := -1
        strNum := checkSubstrNum(line[i:])
        if strNum != -1 {
            newNum = strNum
        } else {
            char := string(line[i])
            charNum, err := strconv.Atoi(char)
            if err != nil {
                continue
            }
            newNum = charNum
        }
        if leftNum == -1 {
            leftNum = newNum
        }  
        rightNum = newNum
    }

    return leftNum, rightNum
}

func findLeftNum(line string) int {
    for i := 0; i < len(line); i++ {
        char := string(line[i])
        charNum, err := strconv.Atoi(char)
        if err != nil {
            continue
        }
        return charNum
    }
    return -1
}

func findRightNum(line string) int {
    for i := len(line) - 1; i >= 0; i-- {
        char := string(line[i])
        charNum, err := strconv.Atoi(char)
        if err != nil {
            continue
        }
        return charNum
    }
    return -1
}

func main() {
    filename := os.Args[1]
    readfile, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
        return
    }

    filescanner := bufio.NewScanner(readfile)
    filescanner.Split(bufio.ScanLines)
    var filelines []string

    for filescanner.Scan() {
        filelines = append(filelines, filescanner.Text())
    }

    readfile.Close()

    sum := 0
    for _, line := range filelines {
        leftNum, rightNum := findNums(line)

        if leftNum == -1 || rightNum == -1 {
            panic("No number :O")
        }

        sum += leftNum * 10 + rightNum
    }

    fmt.Println("Sum", sum)
}
