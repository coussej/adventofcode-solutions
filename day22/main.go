package main

import "fmt"

type player struct {
	hitpoints, mana, damage, remShield, remPoison, remRecharge, manaSpent int
}

var spells map[string]int // {spellname: manacost}

func main() {
	spells = map[string]int{
		"missile":  53,
		"drain":    73,
		"shield":   113,
		"poison":   173,
		"recharge": 229,
	}
	p := player{hitpoints: 50, mana: 500}
	b := player{hitpoints: 51, damage: 9}

	fmt.Println("The smallest amout of mana you can spend while still winning is", p.getSmallestCostForWinningAgainst(b, false))
	fmt.Println("On hard, this amount increase to", p.getSmallestCostForWinningAgainst(b, true))
}

func (p player) getSmallestCostForWinningAgainst(b player, hard bool) (cost int) {

	// If level is hard, player loses 1 hitpoint before anything else happens.
	if hard {
		p.hitpoints--
		if p.hitpoints == 0 {
			return
		}
	}

	// Apply effects before players turn
	p.applyEffects()
	b.applyEffects()

	// Boss could be dead by poison:
	if b.hitpoints <= 0 {
		if p.manaSpent < cost || cost == 0 {
			cost = p.manaSpent
			//fmt.Println("Boss dead by poison:", p, b, cost)
			return
		}
	}

	// Range over the possible spells. If there is enough mana, and casting the
	// spell does not result in a higher cost a the cost of a previously won
	// battle, and casting the spell is allowed (effects that are still ongoing),
	// cast the spell and see where it gets us.
castloop:
	for s, m := range spells {
		pc := p
		bc := b

		// If we don't have enough mana to cast the spell, or casting the spell
		// makes it more expensive then a solution already found, skip.
		if pc.mana < m || (cost > 0 && pc.manaSpent+m >= cost) {
			continue castloop
		}

		// Cast spell by applying the relevant effects to either player or boss.
		switch s {
		case "missile":
			bc.hitpoints -= 4
		case "drain":
			bc.hitpoints -= 2
			pc.hitpoints += 2
		case "shield":
			if pc.remShield > 0 {
				continue castloop
			} else {
				pc.remShield = 6
			}
		case "poison":
			if bc.remPoison > 0 {
				continue castloop
			} else {
				bc.remPoison = 6
			}
		case "recharge":
			if pc.remRecharge > 0 {
				continue castloop
			} else {
				pc.remRecharge = 5
			}
		}

		// Subtract the cost of the spell from the players mana, and add it to the
		// mana spent.
		pc.mana -= m
		pc.manaSpent += m

		// Boss could be dead by spell damage:
		if bc.hitpoints <= 0 {
			if pc.manaSpent < cost || cost == 0 {
				cost = pc.manaSpent
				continue
			}
		}

		// Before applying the boss's damage, apply the effects once more:
		pc.applyEffects()
		bc.applyEffects()

		// Boss could be dead by poison:
		if bc.hitpoints <= 0 {
			if pc.manaSpent < cost || cost == 0 {
				cost = pc.manaSpent
				continue
			}
		}
		// Do boss damage. If the player still has a shield, add the shield value
		// to his hitpoints before subtracting boss's damage.
		if pc.remShield > 0 {
			pc.hitpoints += 7
		}
		pc.hitpoints -= bc.damage

		// Player could be dead. If so, continue to the next spell. If not, both
		// player and boss are still alive, so we must cast another spell. Recurse!
		if pc.hitpoints <= 0 {
			continue
		} else {
			c := pc.getSmallestCostForWinningAgainst(bc, hard)
			if c > 0 && (c < cost || cost == 0) {
				cost = c
			}
		}
	}
	return
}

func (p *player) applyEffects() {
	//   recharge:
	if p.remRecharge > 0 {
		p.mana += 101
		p.remRecharge--
	}
	//   shield
	if p.remShield > 0 {
		p.remShield--
	}
	//   poison
	if p.remPoison > 0 {
		p.remPoison--
		p.hitpoints -= 3
	}
}
