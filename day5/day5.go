package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Mapping struct {
    dst int
    src int
    length int
}

func newMappings(dst int, src int, length int) Mapping {
    return Mapping{
        dst: dst,
        src: src,
        length: length,
    }
}

func initSeeds(line string) []int {
    rhs := strings.Split(line, ": ")[1]
    seedStrs := strings.Split(rhs, " ")

    var seeds []int
    for _, seed := range(seedStrs) {
        seedNum, err := strconv.Atoi(seed)
        if err != nil {
            continue
        }
        seeds = append(seeds, seedNum)
    }

    return seeds
}

func initSeedsPart2(line string) []int {
    rhs := strings.Split(line, ": ")[1]
    seedStrs := strings.Split(rhs, " ")

    var seeds []int
    var seedStart int
    for i, seed := range(seedStrs) {
        if i % 2 == 0 {
            seedNum, err := strconv.Atoi(seed)
            if err != nil {
                panic("NaN")
            }
            seedStart = seedNum
        } else {
            seedRange, err := strconv.Atoi(seed)
            if err != nil {
                panic("NaN")
            }
            for i := seedStart; i < seedStart + seedRange; i++ {
                seeds = append(seeds, i)
            }
        }
    }

    return seeds
}

func initMap(mapping *[]Mapping, line string) {
    mappingStrs := strings.Split(line, " ")
    var mappingNums []int

    for _, str := range(mappingStrs) {
        num, err := strconv.Atoi(str)
        if err != nil {
            continue
        }
        mappingNums = append(mappingNums, num)
    }

    newMapping := newMappings(mappingNums[0], mappingNums[1], mappingNums[2])
    *mapping = append(*mapping, newMapping)
}

func findCorrespondingValue(key int, mappings []Mapping) int {
    for _, mapping := range(mappings) {
        srcStart := mapping.src
        srcEnd := srcStart + mapping.length

        if key >= srcStart && key < srcEnd {
            diff := key - srcStart
            return mapping.dst + diff
        }
    }

    return key
}

func main() {
    filePath := "d5"
    file, _ := os.Open(filePath)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string

    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }

    var seeds                   []int
    var seedToSoil              []Mapping
    var soilToFertilizer        []Mapping
    var fertilizerToWater       []Mapping
    var waterToLight            []Mapping
    var lightToTemperature      []Mapping
    var temperatureToHumidity   []Mapping
    var humidityToLocation      []Mapping

    section := 0
    onInfoLine := false

    for _, line := range(lines) {
        if len(line) == 0 {
            section += 1
            onInfoLine = true
            continue
        }
        switch section {
        case 0:
            seeds = initSeedsPart2(line) 
        case 1:
            if onInfoLine {
                onInfoLine = false
                continue
            } 
            initMap(&seedToSoil, line) 
        case 2: 
            if onInfoLine {
                onInfoLine = false
                continue
            } 
            initMap(&soilToFertilizer, line)
        case 3:
            if onInfoLine {
                onInfoLine = false
                continue
            }
            initMap(&fertilizerToWater, line)
        case 4:
            if onInfoLine {
                onInfoLine = false
                continue
            }
            initMap(&waterToLight, line)
        case 5:
            if onInfoLine {
                onInfoLine = false
                continue
            }
            initMap(&lightToTemperature, line)
        case 6:
            if onInfoLine {
                onInfoLine = false
                continue
            }
            initMap(&temperatureToHumidity, line) 
        case 7:
            if onInfoLine {
                onInfoLine = false
                continue
            }
            initMap(&humidityToLocation, line)
        default:
            break
        }
    }

    // fmt.Println(seeds)
    // fmt.Println(seedToSoil)
    // fmt.Println(soilToFertilizer)
    // fmt.Println(fertilizerToWater)
    // fmt.Println(waterToLight)
    // fmt.Println(lightToTemperature)
    // fmt.Println(temperatureToHumidity)
    // fmt.Println(humidityToLocation)

    var soils []int
    var fertilizers []int
    var waters []int
    var lights []int
    var temperatures []int
    var humidities []int
    var locations []int

    for _, seed := range(seeds) {
        soil := findCorrespondingValue(seed, seedToSoil)
        soils = append(soils, soil)
    }
    // fmt.Println(soils)
    for _, soil := range(soils) {
        fertilizer := findCorrespondingValue(soil, soilToFertilizer)
        fertilizers = append(fertilizers, fertilizer)
    }
    // fmt.Println(fertilizers)
    for _, fertilizer := range(fertilizers) {
        water := findCorrespondingValue(fertilizer, fertilizerToWater)
        waters = append(waters, water)
    }
    // fmt.Println(waters)
    for _, water := range(waters) {
        light := findCorrespondingValue(water, waterToLight)
        lights = append(lights, light)
    } 
    // fmt.Println(lights)
    for _, light := range(lights) {
        temperature := findCorrespondingValue(light, lightToTemperature)
        temperatures = append(temperatures, temperature)
    }
    // fmt.Println(temperatures)
    for _, temperature := range(temperatures) {
        humidity := findCorrespondingValue(temperature, temperatureToHumidity)
        humidities = append(humidities, humidity)
    }
    // fmt.Println(humidities)
    for _, humidity := range(humidities) {
        location := findCorrespondingValue(humidity, humidityToLocation)
        locations = append(locations, location)
    }
    // fmt.Println(locations)

    fmt.Println(slices.Min(locations))
}
