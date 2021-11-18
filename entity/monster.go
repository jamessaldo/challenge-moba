package entity

type Monster struct {
	name   string
	hp     int
	damage int
}

func CreateMonster(name string, hp int, damage int) *Monster {
	return &Monster{name: name, hp: hp, damage: damage}
}

func (monster *Monster) ReceiveDamage(opponent interface{}) {
	switch v := opponent.(type) {
	case *Hero:
		monster.hp = monster.hp - v.damage
	case *Monster:
		monster.hp = monster.hp - v.damage
	}

}

func (monster *Monster) Punch(hitable Hitable) {
	hitable.ReceiveDamage(monster)
}
