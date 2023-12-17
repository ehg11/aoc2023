package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getWinningNumbers(line string) []int {
    lhs := strings.Split(line, "|")[0]
    nums := strings.Split(lhs, ": ")[1]
    trimmedNums := strings.TrimSpace(nums) 
    winningStrs := strings.Split(trimmedNums, " ")

    var winningNums []int
    for _, winningStr := range(winningStrs) {
        winningNumsTrim := strings.TrimSpace(winningStr)
        winningNum, err := strconv.Atoi(winningNumsTrim)
        if err != nil {
            continue
        }
        winningNums = append(winningNums, winningNum)
    }

    return winningNums
} 

func getMyNumbers(line string) []int {
    rhs := strings.Split(line, "|")[1]
    myNumStrings := strings.Split(rhs, " ")

    var myNums []int
    for _, myNumStrings := range(myNumStrings) {
        myNumTrims := strings.TrimSpace(myNumStrings)
        myNum, err := strconv.Atoi(myNumTrims)
        if err != nil {
            continue
        }
        myNums = append(myNums, myNum)
    }

    return myNums
}

func calculatePoints(winningNums []int, myNums []int) int {
    points := 0
    for _, num := range(myNums) {
        if slices.Contains(winningNums, num) {
            if points == 0 {
                points = 1;
            } else {
                points *= 2;
            }
        }
    }

    return points
}

func calculateMatches(winningNums []int, myNums []int) int {
    matches := 0
    for _, num := range(myNums) {
        if slices.Contains(winningNums, num) {
            matches += 1
        }
    }

    return matches
}

func findNumberOfCards(lines []string) []int {
    cards := make([]int, len(lines))
    for i, line := range(lines) {
        cards[i] += 1

        winningNums := getWinningNumbers(line)
        myNums := getMyNumbers(line)

        matches := calculateMatches(winningNums, myNums)

        for j := i + 1; j <= i + matches; j++ {
            cards[j] += cards[i]
        }
    }

    return cards
}

func sum(nums []int) int {
    sum := 0
    for _, num := range(nums) {
        sum += num
    }
    return sum
}

func main() {
    filePath := "d4"
    file, _ := os.Open(filePath)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string

    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }

    cards := findNumberOfCards(lines)
    fmt.Println(cards)
    fmt.Println(sum(cards))
}
