package entity

type Hitable interface {
	ReceiveDamage(x interface{})
}
