package main

import (
	"fmt"

	"github.com/DimitriLaPoudre/InazuScrap/internal/app"
	Player "github.com/DimitriLaPoudre/InazuScrap/internal/player"
)

func main() {
	players := app.Scrap()
	fmt.Println("scrap end")

	players = Player.Sort(players, "sum", true)
	fmt.Println("sort end")

	for _, player := range players {
		// if player.sum == 691 {
		fmt.Println(player)
		// }
	}
}
