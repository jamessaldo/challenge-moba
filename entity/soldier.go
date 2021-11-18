package entity

//SOLDIER
type Soldier struct {
	hp     int
	damage int
}

func CreateSoldier(hp int, damage int) *Soldier {
	return &Soldier{hp: hp, damage: damage}
}

func (soldier *Soldier) ReceiveDamage(opponent interface{}) {
	switch v := opponent.(type) {
	case *Hero:
		soldier.hp = soldier.hp - v.damage
	case *Monster:
		soldier.hp = soldier.hp - v.damage
	}
}
