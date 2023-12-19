package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

type Coord struct {
    row int
    col int
}

func transpose(buffer []string) []string {
    var transposedString []string

    for col := 0; col < len(buffer[0]); col++ {
        var transposedRow []byte
        for row := 0; row < len(buffer); row++ {
            char := buffer[row][col]
            transposedRow = append(transposedRow, char)
        }
        transposedString = append(transposedString, string(transposedRow))
    }

    return transposedString
}

func hasGalaxy(bufline string) bool {
    for _, char := range(bufline) {
        if char == '#' {
            return true
        }
    }
    return false
}

func expandRows(buffer []string) []string {
    var rowExpanded []string

    for _, bufline := range(buffer) {
        rowExpanded = append(rowExpanded, bufline)
        if !hasGalaxy(bufline) {
            rowExpanded = append(rowExpanded, bufline)
        }
    }

    return rowExpanded
}

func expandCols(buffer []string) []string {
    bufferTransposed := transpose(buffer)
    expandRowsTranposed := expandRows(bufferTransposed)
    return transpose(expandRowsTranposed)
}

func findGalaxies(buffer []string) []Coord {
    var galaxies []Coord

    for row, bufline := range(buffer) {
        for col, char := range(bufline) {
            if char == '#' {
                galaxies = append(galaxies, Coord{
                    row: row,
                    col: col,
                })
            }
        }
    }

    return galaxies
}

func distance(lhs Coord, rhs Coord) int {
    xDiff := math.Abs(float64(lhs.row - rhs.row))
    yDiff := math.Abs(float64(lhs.col - rhs.col))

    return int(xDiff + yDiff)
}

func findDistanceSum(galaxies []Coord) int {
    numGalaxies := len(galaxies)
    distanceSum := 0

    for i := 0; i < numGalaxies; i++ {
        for j := i + 1; j < numGalaxies; j++ {
            distanceSum += distance(galaxies[i], galaxies[j]) 
        }
    }

    return distanceSum
}

func show(buffer []string) {
    for _, bufline := range(buffer) {
        fmt.Println(bufline)
    }
}

func findEmptyRows(buffer []string) []int {
    var emptyRows []int
    for row, bufline := range(buffer) {
        if !hasGalaxy(bufline) {
            emptyRows = append(emptyRows, row)
        }
    }

    return emptyRows
}

func findEmptyCols(buffer []string) []int {
    bufferTransposed := transpose(buffer)
    return findEmptyRows(bufferTransposed)
}

func findCrossedEmpty(lhs int, rhs int, empty []int) int {
    numCrossed := 0

    start := lhs
    end := rhs
    if start > end {
        start = rhs
        end = lhs
    }

    for i := start + 1; i < end; i++ {
        if slices.Contains(empty, i) {
            numCrossed += 1
        }
    }

    return numCrossed
}

func findScaledDistanceSum(buffer []string, scale int) int {
    emptyRows := findEmptyRows(buffer)
    emptyCols := findEmptyCols(buffer)
    galaxies := findGalaxies(buffer)
    numGalaxies := len(galaxies)

    distanceSum := 0

    for i := 0; i < numGalaxies; i++ {
        for j := i + 1; j < numGalaxies; j++ {
            lhs := galaxies[i]
            rhs := galaxies[j]
            numCrossedEmptyRows := findCrossedEmpty(lhs.row, rhs.row, emptyRows)
            numCrossedEmptyCols := findCrossedEmpty(lhs.col, rhs.col, emptyCols)


            galaxyDistance := distance(lhs, rhs) + (scale - 1) * (numCrossedEmptyRows + numCrossedEmptyCols)
            distanceSum += galaxyDistance
            // fmt.Println("Comparing", lhs, rhs, "Empty Rows: ", numCrossedEmptyRows, "Empty Cols: ", numCrossedEmptyCols, "Distance: ", galaxyDistance)
        }
    }

    return distanceSum
} 

func main() {
    filename := "d11"
    file, _ := os.Open(filename)

    scanner := bufio.NewScanner(file)
    var buffer []string
    
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    // rowExpanded := expandRows(buffer)
    // expanded := expandCols(rowExpanded)
    //
    // galaxies := findGalaxies(expanded)
    // distanceSum := findDistanceSum(galaxies)
    //
    // fmt.Println(distanceSum)

    scaledDistanceSum := findScaledDistanceSum(buffer, 1000000)
    fmt.Println(scaledDistanceSum)
}
