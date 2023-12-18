package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseLine(history string) []int {
    historySplit := strings.Split(history, " ")
    var historyNums []int
    for _, val := range(historySplit) {
        num, err := strconv.Atoi(val)
        if err != nil {
            panic("paniccccc")
        }
        historyNums = append(historyNums, num)
    } 

    return historyNums
}

func allZeroes(list []int) bool {
    for _, val := range(list) {
        if val != 0 {
            return false;
        }
    }
    return true
}

func extrapolateValue(historyNums []int) int {
    var fullHistory [][]int
    currHistory := historyNums
    fullHistory = append(fullHistory, currHistory)

    for !allZeroes(currHistory) {
        var nextHistory []int
        for i := 0; i < len(currHistory) - 1; i++ {
            diff := currHistory[i + 1] - currHistory[i]
            nextHistory = append(nextHistory, diff)
        }
        currHistory = nextHistory
        fullHistory = append(fullHistory, currHistory)
    }

    for i := len(fullHistory) - 1; i > 0; i-- {
        checkIndex := len(fullHistory[i]) - 1
        top := fullHistory[i - 1][checkIndex]
        bot := fullHistory[i][checkIndex]
        fullHistory[i - 1] = append(fullHistory[i - 1], top + bot)
    }

    return fullHistory[0][len(fullHistory[0]) - 1]
}

func extrapolateValue2(history []int) int {
    var fullHistory [][]int
    currHistory := history
    fullHistory = append(fullHistory, currHistory)

    for !allZeroes(currHistory) {
        var nextHistory []int
        for i := 0; i < len(currHistory) - 1; i++ {
            diff := currHistory[i + 1] - currHistory[i]
            nextHistory = append(nextHistory, diff)
        }
        currHistory = nextHistory
        fullHistory = append(fullHistory, currHistory)
    }

    result := 0
    for i := len(fullHistory) - 1; i > 0; i-- {
        result = fullHistory[i - 1][0] - result
    }

    return result
}

func main() {
    filename := "d9"
    file, _ := os.Open(filename)
    scanner := bufio.NewScanner(file)
    var buffer []string
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    total := 0
    for _, line := range(buffer) {
        lst := parseLine(line) 
        val := extrapolateValue2(lst)

        total += val
    }

    fmt.Println(total)
}
