package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseLine(bufline string) (string, []int) {
    splitBufline := strings.Split(bufline, " ")
    springs := splitBufline[0]
    damagedStrs := splitBufline[1]
    splitDamaged := strings.Split(damagedStrs, ",")
    var damaged []int

    for _, damagedStr := range(splitDamaged) {
        damagedNum, err := strconv.Atoi(damagedStr) 
        if err != nil {
            panic("yeet")
        }
        damaged = append(damaged, damagedNum)
    }

    return springs, damaged
}

func isValid(springs string, damaged []int) bool {
    damagedGroupsAll := strings.Split(springs, ".")
    var damagedGroups []string
    for _, damagedGroup := range(damagedGroupsAll) {
        if len(damagedGroup) > 0 {
            damagedGroups = append(damagedGroups, damagedGroup)
        }
    }

    if len(damagedGroups) != len(damaged) {
        return false
    }

    for i, damagedGroup := range(damagedGroups) {
        if len(damagedGroup) != damaged[i] {
            return false
        }
    }

    return true
}

func replaceChar(str string, replaced byte, index int) string {
    chars := []byte(str)
    chars[index] = replaced
    return string(chars)
}

func countArrangements(springs string, damaged []int, cache *map[string]int) int {
    firstUnknownIndex := strings.Index(springs, "?")
    if firstUnknownIndex == -1 {
        if isValid(springs, damaged) {
            return 1
        } else {
            return 0
        }
    }

    cachedCount, exists := (*cache)[createKey(springs, damaged)]
    if exists {
        return cachedCount
    }

    springBad := replaceChar(springs, '#', firstUnknownIndex)
    springGood := replaceChar(springs, '.', firstUnknownIndex)

    count := countArrangements(springBad, damaged, cache) + countArrangements(springGood, damaged, cache)
    (*cache)[createKey(springs, damaged)] = count
    return count
}

func unfoldedParseLine(bufline string) (string, []int) {
    splitBufline := strings.Split(bufline, " ")
    springs := splitBufline[0]
    damagedStrs := splitBufline[1]

    unfoldedSprings := springs
    unfoldedDamagedStrs := damagedStrs
    for i := 0; i < 4; i++ {
        unfoldedSprings += "?" + springs
        unfoldedDamagedStrs += "," + damagedStrs
    }

    fmt.Println(unfoldedSprings)
    fmt.Println(unfoldedDamagedStrs)

    splitUnfoldedDamaged := strings.Split(unfoldedDamagedStrs, ",")
    var unfoldedDamaged []int
    for _, damagedStr := range(splitUnfoldedDamaged) {
        damagedNum, err := strconv.Atoi(damagedStr) 
        if err != nil {
            panic("yeet")
        }
        unfoldedDamaged = append(unfoldedDamaged, damagedNum)
    }
    
    return unfoldedSprings, unfoldedDamaged
}

func intListToString(nums []int) string {
    str := ""
    for _, num := range(nums) {
        str += string(num) + ","
    }

    return str
}

func createKey(springs string, damaged []int) string {
    return springs + "+" + intListToString(damaged)
}

func countArrangements2(springs string, damaged []int, cache *map[string]int) int {
    if len(springs) == 0 {
        if len(damaged) == 0 {
            return 1
        }
        return 0
    }

    if len(damaged) == 0 {
        if slices.Contains([]byte(springs), '#') {
            return 0
        } 
        return 1
    }

    key := createKey(springs, damaged)
    cachedCount, exists := (*cache)[key]
    if exists {
        return cachedCount
    }

    result := 0
    head := springs[0]
    
    if head == '.' || head == '?' {
        result += countArrangements2(springs[1:], damaged, cache)
    }
    if head == '#' || head == '?' {
        if damaged[0] <= len(springs) {
            if !slices.Contains([]byte(springs[:damaged[0]]), '.') {
                if damaged[0] < len(springs) {
                    if springs[damaged[0]] != '#' {
                        result += countArrangements2(springs[damaged[0] + 1:], damaged[1:], cache)
                    }
                } else if damaged[0] == len(springs) {
                    result += countArrangements2("", damaged[1:], cache)
                }
            }
        }
    }

    (*cache)[key] = result
    return result
}

func main() {
    filename := "d12"
    file, _ := os.Open(filename)
    scanner := bufio.NewScanner(file)
    var buffer []string
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    cache := make(map[string]int)
    totalArragements := 0
    for _, bufline := range(buffer) {
        springs, damaged := unfoldedParseLine(bufline)
        arrangements := countArrangements2(springs, damaged, &cache)
        totalArragements += arrangements
    }

    fmt.Println(totalArragements)
}
