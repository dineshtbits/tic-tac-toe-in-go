package main

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

type Player struct {
	name      string
	pickCount int
}

var attempts [9]*Player
var winningSeq = [][]int{
	{2, 4, 6},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{6, 7, 8},
	{3, 4, 5},
	{0, 1, 2}}

func (p *Player) pick() *Player {
	fmt.Printf("Enter %v your pick: ", p.name)
	var number int
	fmt.Scanf("%d", &number)

	if number > 8 || number < 0 || attempts[number] != nil {
		fmt.Printf("the spot %v is either invalid or already taken", number)
		return p.pick()
	} else {
		p.pickCount++
		attempts[number] = p
		return checkForWinner()
	}
}

func comparePicks(values ...int) bool {
	return attempts[values[0]] != nil &&
		func() bool {
			result := true
			for _, item := range values[1:] {
				if !reflect.DeepEqual(attempts[values[0]], attempts[item]) {
					result = false
					break
				}
			}
			return result
		}()
}

func checkForWinner() *Player {
	for _, seq := range winningSeq {
		if comparePicks(seq...) {
			return attempts[seq[0]]
		}
	}
	return nil
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

func allPicksComplete(players ...*Player) bool {
	var picks int
	for _, player := range players {
		picks += player.pickCount
	}
	return picks >= len(attempts)
}

func createPlayers(count int) []*Player {
	var players []*Player
	for i := 1; i <= count; i++ {
		players = append(players, &Player{name: "P" + strconv.Itoa(i)})
	}
	return players
}

func main() {
	// Initialize program
	noOfPlayers := flag.Int("players", 2, "Please enter the no.of players")
	flag.Parse()
	players := createPlayers(*noOfPlayers)
	gameOver := false

	// Start picks
	for i := 0; !allPicksComplete(players...); {
		showGrid(attempts)
		if players[i].pick() != nil {
			fmt.Printf("%v won\n", players[i].name)
			gameOver = true
			break
		}
		// choosing the next player to pick. Looping over
		if i == len(players)-1 {
			i = 0
		} else {
			i++
		}
	}

	// Result
	if !gameOver {
		fmt.Printf("Match draw\n")
	} else {
		showGrid(attempts)
	}
}
