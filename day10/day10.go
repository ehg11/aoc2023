package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Position struct {
    row int
    col int
    pipe string
}

func getStartPosition(buffer []string) Position {
    for i := 0; i < len(buffer); i++ {
        bufline := []rune(buffer[i])
        for j := 0; j < len(bufline); j++ {
            if bufline[j] == 'S' {
                return Position{
                    row: i,
                    col: j,
                    pipe: "S",
                }
            }
        }
    }

    return Position{
        row: -1,
        col: -1,
        pipe: "!",
    }
}

func getNextPosition(buffer []string, position Position, prevPosition Position) Position {
    currPipe := []rune(position.pipe)[0]

    currRow := position.row
    currCol := position.col
    prevRow := prevPosition.row
    prevCol := prevPosition.col

    nextRow := position.row
    nextCol := position.col

    switch currPipe {
    case '|':
        if (prevRow == currRow - 1) {
            nextRow = currRow + 1
        } else {
            nextRow = currRow - 1
        }
    case '-':
        if (prevCol == currCol - 1) {
            nextCol = currCol + 1
        } else {
            nextCol = currCol - 1
        }
    case 'L':
        if (prevRow == currRow - 1) {
            nextCol = currCol + 1
        } else {
            nextRow = currRow - 1
        }
    case 'J':
        if (prevRow == currRow - 1) {
            nextCol = currCol - 1
        } else {
            nextRow = currRow - 1
        }
    case '7':
        if (prevRow == currRow + 1) {
            nextCol = currCol - 1
        } else {
            nextRow = currRow + 1
        }
    case 'F':
        if (prevRow == currRow + 1) {
            nextCol = currCol + 1
        } else {
            nextRow = currRow + 1
        }
    case 'S':
        totalRows := len(buffer)
        totalCols := len(buffer[0])
        if (currRow > 0) {
            nextChar := buffer[currRow - 1][currCol]
            if  nextChar == '|' ||
                nextChar == '7' ||
                nextChar == 'F' {
                nextRow = currRow - 1 
                break
            }
        }
        if (currRow < totalRows - 1) {
            nextChar := buffer[currRow + 1][currCol]
            if  nextChar == '|' ||
                nextChar == 'L' ||
                nextChar == 'J' {
                nextRow = currRow + 1
                break
            }
        }
        if (currCol > 0) {
            nextChar := buffer[currRow][currCol - 1]
            if  nextChar == '-' ||
                nextChar == 'L' ||
                nextChar == 'F' {
                nextCol = currCol - 1
                break
            }
        }
        if (currCol < totalCols - 1) {
            nextChar := buffer[currRow][currCol + 1]
            if  nextChar == '-' ||
                nextChar == 'J' ||
                nextChar == '7' {
                nextCol = currCol + 1
                break
            }
        }
    }

    return Position{
        row: nextRow,
        col: nextCol,
        pipe: string(buffer[nextRow][nextCol]),
    }
}

func samePosition(lhs Position, rhs Position) bool {
    return lhs.row == rhs.row && lhs.col == rhs.col
}

func getPath(buffer []string) []Position {
    startPosition := getStartPosition(buffer)
    nextPosition := getNextPosition(buffer, startPosition, Position{})

    var path []Position
    path = append(path, startPosition)

    for !samePosition(startPosition, nextPosition) {
        prevPosition := path[len(path) - 1]
        path = append(path, nextPosition)
        nextPosition = getNextPosition(buffer, nextPosition, prevPosition)
    }

    return path
}

func getStartChar(buffer []string) string {
    startPosition := getStartPosition(buffer)
    
    n := '.'
    e := '.'
    s := '.'
    w := '.'
    if startPosition.row > 0 {
        n = rune(buffer[startPosition.row - 1][startPosition.col])
    }
    if startPosition.col < len(buffer[0]) {
        e = rune(buffer[startPosition.row][startPosition.col + 1])
    }
    if startPosition.row < len(buffer) {
        s = rune(buffer[startPosition.row + 1][startPosition.col])
    }
    if startPosition.col > 0 {
        w = rune(buffer[startPosition.row][startPosition.col - 1])
    }

    if n == '|' || n == '7' || n == 'F' {
        if s == '|' || s == 'L' || s == 'J' {
            return "|"
        }
        if e == '-' || e == '7' || e == 'J' {
            return "L"
        }
        if w == '-' || w == 'L' || w == 'F' {
            return "J"
        }
    }
    if e == '-' || e == '7' || e == 'J' {
        if w == '-' || w == 'L' || w == 'F' {
            return "-"
        }
        if s == '|' || s == 'L' || s == 'J' {
            return "F"
        }
    }
    if s == '|' || s == 'L' || s == 'J' {
        if w == '-' || w == 'L' || w == 'F' {
            return "7"
        }
    }

    return "."
}

func isInside(position Position, buffer []string) bool {
    currRow := position.row
    currCol := position.col
    totalCols := len(buffer[0])

    intersections := 0
    unclearIntersections := []byte { 'F', '7', 'L', 'J' }
    unclearTops := []byte { 'F', '7' }
    unclearBots := []byte {'L', 'J' }
    var seenUnclears []byte
    for i := currCol + 1; i < totalCols; i++ {
        if buffer[currRow][i] == '|' {
            intersections += 1
        }
        if slices.Contains(unclearIntersections, buffer[currRow][i]) {
            seenUnclears = append(seenUnclears, buffer[currRow][i])
        }
        if len(seenUnclears) == 2 {
            inTop := slices.Contains(unclearTops, seenUnclears[0]) || slices.Contains(unclearTops, seenUnclears[1])
            inBot := slices.Contains(unclearBots, seenUnclears[0]) || slices.Contains(unclearBots, seenUnclears[1])

            if inTop && inBot {
                intersections += 1
            } else if inTop || inBot {
                intersections += 2
            } else {
                intersections += 0
            }

            seenUnclears = nil
        }
    }

    return intersections % 2 == 1
}

func getCleanRow(row int, bufline string, path []Position) string {
    cleanBufline := []byte(bufline)

    for col, char := range(cleanBufline) {
        if char == '.' {
            continue
        }
        onPath := false
        for _, pos := range(path) {
            if row == pos.row && col == pos.col {
                onPath = true
                break
            }
        }
        if !onPath {
            cleanBufline[col] = '.'
        }
    }

    return string(cleanBufline)
}

func countInside(buffer []string) int {
    numInside := 0

    startPosition := getStartPosition(buffer)
    startChar := getStartChar(buffer)

    path := getPath(buffer)
    for row, bufline := range(buffer) {
        cleanRow := getCleanRow(row, bufline, path)
        buffer[row] = cleanRow
    }

    startRow := buffer[startPosition.row]
    modifiedStartRow := []rune(startRow)
    modifiedStartRow[startPosition.col] = []rune(startChar)[0]
    buffer[startPosition.row] = string(modifiedStartRow)

    for i, row := range(buffer) {
        for j, char := range(row) {
            if char == '.' {
                position := Position{
                    row: i,
                    col: j,
                    pipe: ".",
                }
                inside := isInside(position, buffer)
                if inside {
                    numInside += 1
                }
            }
        }
    }

    return numInside
}

func main() {
    filename := "d10"
    file, _ := os.Open(filename)
    scanner := bufio.NewScanner(file)
    var buffer []string
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    numInside := countInside(buffer)
    fmt.Println(numInside)
}
