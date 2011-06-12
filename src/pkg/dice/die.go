package dice

import (
	"rand"
	)

type Die uint8

func RollDie() Die{
	return Die(rand.Intn(6) + 1)
}

func (d Die) PointValue() int{
	switch d {
		case 5:
			return 50
		case 1:
			return 100
	}
	return 0	
}

func (d Die) TriplePointValue() int{
	if d == 1{
		return 1000
	}
	return int(d)*100
}
