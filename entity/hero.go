package entity

//HERO
type Hero struct {
	name   string
	hp     int
	damage int
}

func CreateHero(name string, hp int, damage int) *Hero {
	return &Hero{name: name, hp: hp, damage: damage}
}

func (hero *Hero) ReceiveDamage(opponent interface{}) {
	switch v := opponent.(type) {
	case *Hero:
		hero.hp = hero.hp - v.damage
	case *Monster:
		hero.hp = hero.hp - v.damage
	}
}

func (hero *Hero) Punch(hitable Hitable) {
	hitable.ReceiveDamage(hero)
}

// func (hero *Hero) ReceiveDamageMonster(hitter *Monster) {
// 	hero.hp = hero.hp - hitter.damage
// }
