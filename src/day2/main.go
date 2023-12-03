package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("PartOne: %d\n", partOne(string(contents)))
}

type Color string

func parseColor(in string) Color {
	switch in {
	case string(blue):
		return blue
	case string(green):
		return green
	case string(red):
		return red
	default:
		log.Fatalf("invalid color! %s", in)
		return undefined
	}
}

const (
	blue      Color = "blue"
	green     Color = "green"
	red       Color = "red"
	undefined Color = "undefined"
)

type game struct {
	limits   map[Color]int
	currents map[Color]int
}

// isPossible checks if the current draws are possible given the loaded bag.
func (g *game) isPossible(color Color, drawCount int) bool {
	if cnt, ok := g.currents[color]; ok {
		g.currents[color] = cnt + drawCount
	} else {
		g.currents[color] = drawCount
	}
	if g.limits[color] < drawCount {
		return false
	}
	return true
}

func (g *game) reset() {
	g.currents = map[Color]int{}
}

func partOne(input string) int {
	g := &game{
		limits:   map[Color]int{blue: 14, green: 13, red: 12},
		currents: map[Color]int{},
	}
	fmt.Printf("Loaded game with the following dice: %v\n", g)

	validGameIDs := []int{}
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		lineParts := strings.Split(line, ":")
		gameID := parseGameId(lineParts[0])

		//  Parse out the game cubes
		handOK := true

		hands := strings.Split(strings.TrimSpace(lineParts[1]), ";")
		for _, hand := range hands {
			diceGroups := strings.Split(strings.TrimSpace(hand), ",")
			for _, diceCount := range diceGroups {
				diceCount := strings.TrimSpace(diceCount)
				diceCountParts := strings.Split(diceCount, " ")
				count, err := strconv.Atoi(diceCountParts[0])
				if err != nil {
					log.Fatalf("failed to parse dice count: %v", err)
				}
				col := diceCountParts[1]
				if !g.isPossible(parseColor(col), count) {
					handOK = false
				}
			}
			g.reset()
		}
		if handOK {
			validGameIDs = append(validGameIDs, gameID)
		}
	}
	ans := 0
	for _, id := range validGameIDs {
		ans = ans + id
	}
	return ans
}

func parseGameId(gameStmt string) int {
	stmts := strings.Split(gameStmt, " ")
	gameID, err := strconv.Atoi(stmts[1])
	if err != nil {
		log.Fatalf("failed to convert %q, %v", gameStmt, err)
	}
	return gameID
}
