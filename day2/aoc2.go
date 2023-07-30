package main

import (
	//"fmt"
	//"strconv"

	"fmt"
	"strings"

	"golang.design/x/clipboard"
)

func main() {

	err := clipboard.Init()
	if err != nil {
		panic(err)
	}
	SolveDay2_part1()
	SolveDay2_part2()
}

func SolveDay2_part1() {
	const rock, paper, scissors, loss, draw, win int = 1, 2, 3, 0, 3, 6 // rock, paper, scissors, loss, draw, win r, p, s, l, d ,w
	var score int
	content := clipboard.Read(clipboard.FmtText)
	//var keys []string
	if content == nil {
		panic("No clipboard content")
	}
	sets := string(content)
	setsMap := strings.Split(sets, "\r\n")
	for _, vals := range setsMap {
		if len(vals) == 0 { // Skip empty lines
			continue
		}
		if vals[0] == 'A' && vals[2] == 'X' { //Rock to Rock
			score += rock + draw
		}
		if vals[0] == 'A' && vals[2] == 'Y' { //Rock to Paper
			score += paper + win
		}
		if vals[0] == 'A' && vals[2] == 'Z' { //Rock to Scissors
			score += scissors + loss
		}
		if vals[0] == 'B' && vals[2] == 'X' { // Paper to Rock
			score += rock + loss
		}
		if vals[0] == 'B' && vals[2] == 'Y' { // Paper to Paper
			score += paper + draw
		}
		if vals[0] == 'B' && vals[2] == 'Z' { // Paper to Scissors
			score += scissors + win
		}
		if vals[0] == 'C' && vals[2] == 'X' { // Scissors to Rock
			score += rock + win
		}
		if vals[0] == 'C' && vals[2] == 'Y' { // Scissors to Paper
			score += paper + loss
		}
		if vals[0] == 'C' && vals[2] == 'Z' { // Scissors to Scissors
			score += scissors + draw
		}

	}
	fmt.Println("If my brain works right I should get: ", score)
	//fmt.Println(setsMap)
}

func SolveDay2_part2() {
	const rock, paper, scissors, loss, draw, win int = 1, 2, 3, 0, 3, 6 // rock, paper, scissors, loss, draw, win r, p, s, l, d ,w
	var score int
	content := clipboard.Read(clipboard.FmtText)
	//var keys []string
	if content == nil {
		panic("No clipboard content")
	}
	sets := string(content)
	setsMap := strings.Split(sets, "\r\n")
	for _, vals := range setsMap {
		if len(vals) == 0 { // Skip empty lines
			continue
		}
		if vals[0] == 'A' && vals[2] == 'X' { //Rock to Loss
			score += scissors + loss
		}
		if vals[0] == 'A' && vals[2] == 'Y' { //Rock to Draw
			score += rock + draw
		}
		if vals[0] == 'A' && vals[2] == 'Z' { //Rock to Win
			score += paper + win
		}
		if vals[0] == 'B' && vals[2] == 'X' { // Paper to Loss
			score += rock + loss
		}
		if vals[0] == 'B' && vals[2] == 'Y' { // Paper to Draw
			score += paper + draw
		}
		if vals[0] == 'B' && vals[2] == 'Z' { // Paper to Win
			score += scissors + win
		}
		if vals[0] == 'C' && vals[2] == 'X' { // Scissors to Loss
			score += paper + loss
		}
		if vals[0] == 'C' && vals[2] == 'Y' { // Scissors to Draw
			score += scissors + draw
		}
		if vals[0] == 'C' && vals[2] == 'Z' { // Scissors to Win
			score += rock + win
		}

	}
	fmt.Println("If my brain works right I should get pt2: ", score)
	//fmt.Println(setsMap)
}
