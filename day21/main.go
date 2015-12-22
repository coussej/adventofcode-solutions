package main

import "fmt"

type player struct {
	hitpoints, damage, armor int
	items                    []item
}

type item struct {
	name                string
	cost, damage, armor int
}

func main() {
	p := player{100, 0, 0, []item{}}
	boss := player{103, 9, 2, []item{}}

	minCostWin := 0
	maxCostLoss := 0

	for _, pl := range p.getAllPossiblePlayers() {
		win := pl.fight(boss)
		if win && (pl.totalCost() < minCostWin || minCostWin == 0) {
			minCostWin = pl.totalCost()
		}
		if !win && pl.totalCost() > maxCostLoss {
			maxCostLoss = pl.totalCost()
		}
	}
	fmt.Println("You can win while paying only", minCostWin, "...")
	fmt.Println("... but you can still lose when paying", maxCostLoss, "!")
}

func (p1 player) fight(p2 player) (wins bool) {
	p1h := p1.hitpoints
	p2h := p2.hitpoints
	for {
		p2h -= max(p1.totalDamage()-p2.totalArmor(), 1)
		if p2h <= 0 {
			return true
		}
		p1h -= max(p2.totalDamage()-p1.totalArmor(), 1)
		if p1h <= 0 {
			return false
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (p player) getAllPossiblePlayers() (ps []player) {
	// make inventory
	weapons := []item{
		item{"Dagger", 8, 4, 0},
		item{"Shortsword", 10, 5, 0},
		item{"Warhammer", 25, 6, 0},
		item{"Longsword", 40, 7, 0},
		item{"Greataxe", 74, 8, 0},
	}
	armor := []item{
		item{"Leather", 13, 0, 1},
		item{"Chainmail", 31, 0, 2},
		item{"Splintmail", 53, 0, 3},
		item{"Bandedmail", 75, 0, 4},
		item{"Platemail", 102, 0, 5},
	}
	rings := []item{
		item{"Damage +1", 25, 1, 0},
		item{"Damage +2", 50, 2, 0},
		item{"Damage +3", 100, 3, 0},
		item{"Defense +1", 20, 0, 1},
		item{"Defense +2", 40, 0, 2},
		item{"Defense +3", 80, 0, 3},
	}
	// create template player
	ps = []player{p}
	// loop weapons
	for _, w := range weapons {
		// add a player with each weapon, no armor or rings.
		pw := p.copy()
		pw.items = append(pw.items, w)
		ps = append(ps, pw)
		// loop armor
		for _, a := range armor {
			pa := pw.copy()
			pa.items = append(pa.items, a)
			ps = append(ps, pa)
			// loop rings
			for i, r := range rings {
				pr := pa.copy()
				pr.items = append(pr.items, r)
				ps = append(ps, pr)
				// add another ring
				for i2, r2 := range rings {
					if i2 != i {
						pr2 := pr.copy()
						copy(pr2.items, pr.items)
						pr2.items = append(pr2.items, r2)
						ps = append(ps, pr2)
					}
				}
			}
		}
	}
	return
}

func (p player) copy() (pl player) {
	pl = player{p.hitpoints, p.damage, p.armor, []item{}}
	for _, i := range p.items {
		pl.items = append(pl.items, i)
	}
	return
}

func (p player) totalDamage() (damage int) {
	damage = p.damage
	for _, i := range p.items {
		damage += i.damage
	}
	return
}

func (p player) totalArmor() (armor int) {
	armor = p.armor
	for _, i := range p.items {
		armor += i.armor
	}
	return
}

func (p player) totalCost() (cost int) {
	for _, i := range p.items {
		cost += i.cost
	}
	return
}
