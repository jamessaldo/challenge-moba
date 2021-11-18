package main

import (
	"fmt"

	. "enigma.com/challenge-moba/entity"
)

func main() {
	var saitama *Hero = CreateHero("saitama", 1000, 400)
	var thor *Hero = CreateHero("thor", 1000, 400)
	var gorgon *Monster = CreateMonster("gorgon", 1000, 400)
	var ryan *Soldier = CreateSoldier(1000, 400)

	fmt.Println("Before Punched")
	fmt.Println(*saitama)
	fmt.Println(*thor)
	fmt.Println(*gorgon)
	fmt.Println(*ryan)

	saitama.Punch(thor)
	saitama.Punch(gorgon)
	saitama.Punch(ryan)

	fmt.Printf("\nAfter Punched by saitama\n")
	fmt.Println(*saitama)
	fmt.Println(*thor)
	fmt.Println(*gorgon)
	fmt.Println(*ryan)

	gorgon.Punch(saitama)
	gorgon.Punch(thor)
	gorgon.Punch(ryan)

	fmt.Printf("\nAfter Punched by gorgon\n")
	fmt.Println(*saitama)
	fmt.Println(*thor)
	fmt.Println(*gorgon)
	fmt.Println(*ryan)
}
