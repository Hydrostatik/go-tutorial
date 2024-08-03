package businesslayer

type Character struct {
	Name       string
	Profession string
	Health     int
}

func (x Character) Attack(target *Character) int {
	if (Die{}.Roll1To(20)) > 11 {
		attackRoll := Die{}.Roll1To(6)
		target.Health -= attackRoll
		return attackRoll
	}

	return 0
}
