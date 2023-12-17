package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type CamelCards struct {
    hand string
    bid int
    variant int
}

func newHandAndBid(hand string, bid int) CamelCards {
    return CamelCards{
        hand: hand,
        bid: bid,
        variant: -1,
    }
}

func getCardCounts(hand string) map[string]int {
    cardCountMap := make(map[string]int)
    chars := []rune(hand)

    for _, char := range(chars) {
        cardCountMap[string(char)]++
    }

    return cardCountMap
}

func parseLine(bufline string) CamelCards {
    split := strings.Split(bufline, " ")
    hand := split[0]
    bid, err := strconv.Atoi(split[1])
    if err != nil {
        panic("goofed")
    }

    return newHandAndBid(hand, bid)
}

func fiveOfAKind(cardCounts map[string]int) bool {
    for _, v := range(cardCounts) {
        if v == 5 {
            return true
        }
    }
    return false
}

func fourOfAKind(cardCounts map[string]int) bool {
    for _, v := range(cardCounts) {
        if v == 4 {
            return true
        }
    }
    return false
}

func fullHouse(cardCounts map[string]int) bool {
    hasTriple := false
    hasDouble := false
    for _, v := range(cardCounts) {
        if v == 3 {
            hasTriple = true
        }
        if v == 2 {
            hasDouble = true
        }
    }
    return hasDouble && hasTriple
}

func threeOfAKind(cardCounts map[string]int) bool {
    for _, v := range(cardCounts) {
        if v == 3 {
            return true
        }
    }
    return false
}

func twoPair(cardCounts map[string]int) bool {
    numPairs := 0
    for _, v := range(cardCounts) {
        if v == 2 {
            numPairs += 1
        }
    }
    return numPairs == 2
}

func onePair(cardCounts map[string]int) bool {
    numPairs := 0
    for _, v := range(cardCounts) {
        if v == 2 {
            numPairs += 1
        }
    }
    return numPairs == 1
}

func highCard(cardCounts map[string]int) bool {
    numSingles := 0
    for _, v := range(cardCounts) {
        if v == 1 {
            numSingles += 1
        }
    }
    return numSingles == 5
}


func identifyVariant(camelCard CamelCards) int {
    cardCounts := getCardCounts(camelCard.hand)
    if (fiveOfAKind(cardCounts)) {
        // fmt.Println("fiveOfAKind")
        camelCard.variant = 7
    } else if (fourOfAKind(cardCounts)) {
        // fmt.Println("fourOfAKind")
        camelCard.variant = 6
    } else if (fullHouse(cardCounts)) {
        // fmt.Println("fullHouse")
        camelCard.variant = 5
    } else if (threeOfAKind(cardCounts)) {
        // fmt.Println("threeOfAKind")
        camelCard.variant = 4
    } else if (twoPair(cardCounts)) {
        // fmt.Println("twoPair")
        camelCard.variant = 3
    } else if (onePair(cardCounts)) {
        // fmt.Println("onePair")
        camelCard.variant = 2
    } else if (highCard(cardCounts)) {
        // fmt.Println("highCard")
        camelCard.variant = 1
    } 
    return camelCard.variant
}

func findMaximumPossibleVariant(camelCard CamelCards) int {
    chars := []rune(camelCard.hand)
    jokerIndex := slices.Index(chars, 'J')

    if jokerIndex == -1 {
        return identifyVariant(camelCard)
    }

    cardTypes := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
    maximumVariant := -1
    for _, cardType := range(cardTypes) {
        chars[jokerIndex] = cardType
        camelCard.hand = string(chars)
        newHandVariant := findMaximumPossibleVariant(camelCard)
        if newHandVariant > maximumVariant {
            maximumVariant = newHandVariant
        }
    }

    return maximumVariant
}

func main() {
    filepath := "d7"
    file, _ := os.Open(filepath)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var buffer []string
    for scanner.Scan() {
        buffer = append(buffer, scanner.Text())
    }

    var handAndBids []CamelCards
    for _, bufline := range(buffer) {
        handAndBids = append(handAndBids, parseLine(bufline))
    }

    for i, handAndBid := range(handAndBids) {
        handAndBids[i].variant = findMaximumPossibleVariant(handAndBid)
    }

    cardStrength := map[string]int{
        "A": 1,
        "K": 2,
        "Q": 3,
        "T": 4,
        "9": 5,
        "8": 6,
        "7": 7,
        "6": 8,
        "5": 9,
        "4": 10,
        "3": 11,
        "2": 12,
        "J": 13,
    }

    sort.Slice(handAndBids, func(i int, j int) bool {
        lhs := handAndBids[i]
        rhs := handAndBids[j]

        if lhs.variant < rhs.variant {
            return true
        } else if lhs.variant > rhs.variant {
            return false
        } else {
            for i := 0; i < 5; i++ {
                lhsStrength := cardStrength[string(lhs.hand[i])]
                rhsStrength := cardStrength[string(rhs.hand[i])]
                if lhsStrength != rhsStrength {
                    return lhsStrength > rhsStrength
                }
            }
        }
        return true
    })

    totalWinnings := 0
    for i, handAndBid := range(handAndBids) {
        totalWinnings += (i + 1) * handAndBid.bid
    }

    fmt.Println(totalWinnings)
}
