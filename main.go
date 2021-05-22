package main

import (
	"fmt"
)

type Player struct {
	id   int
	name string
}

var attempts [9]*Player

func (p *Player) pick() int {
	fmt.Printf("Enter %v your pick: ", p.name)
	var number int
	fmt.Scanf("%d", &number)

	if number > 8 || number < 0 || attempts[number] != nil {
		fmt.Printf("the spot %v is either invalid or already taken, available spots %v try again.....", number, attempts)
		return p.pick()
	} else {
		attempts[number] = p
		return checkForWinningSeq(attempts)
	}
}

func checkForWinningSeq(attempts [9]*Player) int {
	sums := []int{0, 0, 0, 0, 0, 0, 0, 0}
	sums[0] = getId(attempts[2]) + getId(attempts[4]) + getId(attempts[6])
	sums[1] = getId(attempts[0]) + getId(attempts[3]) + getId(attempts[6])
	sums[2] = getId(attempts[1]) + getId(attempts[4]) + getId(attempts[7])
	sums[3] = getId(attempts[2]) + getId(attempts[5]) + getId(attempts[8])
	sums[4] = getId(attempts[0]) + getId(attempts[4]) + getId(attempts[8])
	sums[5] = getId(attempts[6]) + getId(attempts[7]) + getId(attempts[8])
	sums[6] = getId(attempts[3]) + getId(attempts[4]) + getId(attempts[5])
	sums[7] = getId(attempts[0]) + getId(attempts[1]) + getId(attempts[2])

	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0
}

func getId(p *Player) int {
	if p == nil {
		return 0
	} else {
		return p.id
	}
}

func showGrid(attempts [9]*Player) {
	fmt.Printf("Board: \n")
	for i := 0; i <= 8; i++ {
		if attempts[i] != nil {
			fmt.Printf("|	%v	|", attempts[i].name)
		} else {
			fmt.Printf("|	%v	|", i)
		}
		if (i+1)%3 == 0 {
			fmt.Printf("\n")
			fmt.Printf("-------------------------------------------------")
			fmt.Printf("\n")
		}
	}
}

func main() {

	player1 := &Player{name: "P1", id: 1}
	player2 := &Player{name: "P2", id: 10}

	gameOver := false

	for !gameOver {
		showGrid(attempts)
		if player1.pick() == 1 {
			fmt.Printf("%v won\n", player1.name)
			gameOver = true
			break
		}
		showGrid(attempts)
		if player2.pick() == 2 {
			fmt.Printf("%v won\n", player2.name)
			gameOver = true
			break
		}
	}

	if !gameOver {
		fmt.Printf("Match draw\n")
	} else {
		showGrid(attempts)
	}
}
