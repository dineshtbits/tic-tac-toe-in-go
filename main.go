package main

import (
	"fmt"
)

type Player struct {
	name   string
	picks  []int
	symbol string
}

var attempts []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var winningSeq [][]int = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 4, 7}, {3, 6, 9}, {1, 5, 9}, {3, 5, 7}, {2, 5, 8}}

func (p *Player) pick() bool {
	if len(attempts) <= 0 {
		return false
	}

	fmt.Printf("Enter %v your pick: ", p.name)
	var number int
	fmt.Scanf("%d", &number)

	if indexOf(attempts, number) == -1 {
		fmt.Printf("the spot %v is either invalid or already taken, available spots %v try again.....", number, attempts)
		return p.pick()
	} else {
		// add to player picks
		p.picks = append(p.picks, number)
		// remove from attempts
		attempts = append(attempts[:indexOf(attempts, number)], attempts[indexOf(attempts, number)+1:]...)
		// Check for winning sequence
		return checkForWinningSeq(p.picks)
	}
}

func checkForWinningSeq(picks []int) bool {
	matched := false
	for _, seq := range winningSeq {
		if sequencePicked(seq, picks) {
			matched = true
			break
		}
	}
	return matched
}

func sequencePicked(seq []int, picks []int) bool {
	found := true
	for _, ele := range seq {
		if indexOf(picks, ele) == -1 {
			found = false
			break
		}
	}
	return found
}

func indexOf(items []int, element int) int {
	index := -1
	for i, val := range items {
		if val == element {
			index = i
		}
	}
	return index
}

func showGrid(player1 *Player, player2 *Player) {
	fmt.Printf("Board: \n")
	for i := 1; i <= 9; i++ {
		if indexOf(player1.picks, i) >= 0 {
			fmt.Printf("|	%v	|", player1.name)
		} else if indexOf(player2.picks, i) >= 0 {
			fmt.Printf("|	%v	|", player2.name)
		} else {
			fmt.Printf("|	%v	|", i)
		}

		if i%3 == 0 {
			fmt.Printf("\n")
			fmt.Printf("---------------------------------------------------")
			fmt.Printf("\n")
		}
	}
}

func main() {

	player1 := &Player{name: "dinesh"}
	player2 := &Player{name: "karthik"}

	for len(attempts) > 0 {
		showGrid(player1, player2)
		if player1.pick() {
			fmt.Printf("%v won\n", player1.name)
			break
		}

		showGrid(player1, player2)
		if player2.pick() {
			fmt.Printf("%v won\n", player2.name)
			break
		}
	}
	if len(attempts) <= 0 {
		fmt.Printf("Match draw\n")
	}
}
