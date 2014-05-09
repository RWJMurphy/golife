package main

import (
	"bufio"
	"fmt"
	"os"
	"rea/hackday/16/conway/libconway"
	"time"
)

func main() {
	var width, height int

	fmt.Scanf("%d,%d", &width, &height)
	game := conway.NewGame(width, height)

	stdin := bufio.NewReader(os.Stdin)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if rune, _, _ := stdin.ReadRune(); rune == '*' {
				game.Current.SetCellAlive(x, y)
			}
		}
		stdin.ReadRune()
	}

	for {
		game.Print()
		game.Next()
		if game.Current.Equal(game.Prev) {
			fmt.Println("Equilibrium.")
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	return
}
