package main

import "fmt"

type Action interface {
	Attack()
	Cast()
}

type Mage struct {
	Mana int
	Skills []SkillSet
}

type SkillSet struct {
	Skill string
	ManaUse uint
}

type Knight struct {
	Str int
}

func (k Knight) Attack() {
	fmt.Printf("Knight Attack: %v\n", k.Str)
}

func (m Mage) Cast() {
	fmt.Printf("Mage Cast: %v\n", m.Mana)
}

func game() {
	draken := Knight{Str: 20}
	draken.Attack()

	juvin := Mage{
		Mana: 100,
		Skills: []SkillSet{
			{Skill: "Fireball", ManaUse: 10},
			{Skill: "Icebolt", ManaUse: 5},
		},
	}
	juvin.Cast()
}
