package main

import "fmt"

type Action interface {
	attack()
}

type Character struct {
	name string
	level uint64
	hp uint
	sp uint
	atk uint
	def uint
	int uint
}

func (character *Character) attack(target *Character) {
	damage := character.atk - target.def
	target.hp -= damage
}

func main() {
	hero := &Character{
		name: "Roto",
		level: 1,
		hp: 100,
		sp: 100,
		atk: 20,
		def: 8,
		int: 3,
	}

	enemy := &Character{
		name: "Jvilla",
		level: 1,
		hp: 120,
		sp: 100,
		atk: 14,
		def: 10,
		int: 5,
	}

	fmt.Printf("Enemy HP: %d\n", enemy.hp)
	hero.attack(enemy)
	fmt.Printf("Enemy HP: %d\n", enemy.hp)
}
