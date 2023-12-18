package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Destinations struct {
    left string
    right string
}

func parseLines(buffer []string) map[string]Destinations {
    destMap := make(map[string]Destinations)
    for _, bufline := range(buffer) {
        buflineSplit := strings.Split(bufline, " = ")
        key := buflineSplit[0]
        dests := string([]rune(buflineSplit[1])[1:len(buflineSplit[1])-1])
        destsSplit := strings.Split(dests, ", ")
        destStruct := Destinations{
            left: destsSplit[0],
            right: destsSplit[1],
        }

        destMap[key] = destStruct
    }

    return destMap
}

func parseLines2(buffer []string) (map[string]Destinations, []string) {
    destMap := make(map[string]Destinations)
    var startPositions []string
    for _, bufline := range(buffer) {
        buflineSplit := strings.Split(bufline, " = ")
        key := buflineSplit[0]
        dests := string([]rune(buflineSplit[1])[1:len(buflineSplit[1])-1])
        destsSplit := strings.Split(dests, ", ")
        destStruct := Destinations{
            left: destsSplit[0],
            right: destsSplit[1],
        }

        destMap[key] = destStruct

        if []rune(key)[len(key) - 1] == 'A' {
            startPositions = append(startPositions, key)
        }
    }

    return destMap, startPositions
}

func getNumSteps(steps string, destMap map[string]Destinations) int {
    curr := "AAA"
    numSteps := 0
    stepsArray := []rune(steps)
    for true {
        if curr == "ZZZ" {
            break
        }
        index := numSteps % len(stepsArray) 
        direction := stepsArray[index]
        if (direction == 'L') {
            curr = destMap[curr].left
        } else {
            curr = destMap[curr].right
        }
        numSteps += 1
    }

    return numSteps
}

func allOnZ(positions []string) bool {
    for _, pos := range(positions) {
        posArray := []rune(pos)
        if posArray[len(pos) - 1] != 'Z' {
            return false
        } 
    }
    return true
}

func getNumSteps2(steps string, destMap map[string]Destinations, startPositions []string) int {
    numSteps := 0
    stepsArray := []rune(steps)
    positions := startPositions
    for true {
        fmt.Println(positions, numSteps)
        if allOnZ(positions) {
            break
        }
        index := numSteps % len(stepsArray)
        direction := stepsArray[index]
        if direction == 'L' {
            for i, pos := range(positions) {
                positions[i] = destMap[pos].left
            }
        } else {
            for i, pos := range(positions) {
                positions[i] = destMap[pos].right
            }
        }
        numSteps += 1
    }

    return numSteps
}

func getNumSteps3(steps string, destMap map[string]Destinations, startPosition string) int {
    numSteps := 0
    stepsArray := []rune(steps)
    curr := startPosition;
    for true {
        if []rune(curr)[len(curr) - 1] == 'Z' {
            break
        }
        index := numSteps % len(stepsArray) 
        direction := stepsArray[index]
        if (direction == 'L') {
            curr = destMap[curr].left
        } else {
            curr = destMap[curr].right
        }
        numSteps += 1
    }

    return numSteps
}

func gcd(a int, b int) int {
    for b != 0 {
        temp := b
        b = a % b
        a = temp
    }

    return a
}

func lcm(a int, b int) int {
    return (a * b) / gcd(a, b)
}

func lcmOfSlice(slice []int) int {
    result := slice[0]

    for _, num := range(slice) {
        result = lcm(result, num)
    }

    return result
}

func main() {
    filename := "d8"
    file, _ := os.Open(filename)
    scanner := bufio.NewScanner(file)

    var buffer []string
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    steps := buffer[0]
    buffer = buffer[2:]

    destMap, startPositions := parseLines2(buffer)
    var allNumSteps []int
    for _, startPos := range(startPositions) {
        numSteps := getNumSteps3(steps, destMap, startPos)
        allNumSteps = append(allNumSteps, numSteps)
    }

    fmt.Println(allNumSteps)
    numSteps := lcmOfSlice(allNumSteps)
    fmt.Println(numSteps)
}
