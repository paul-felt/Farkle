package strategies

import "dice"

/* 
 * Keep the very first roll, regardless of the point value 
 */
type FirstRoll int

func NewFirstRoll() FirstRoll{
	return FirstRoll(0)
}

func (s FirstRoll) Keep(score int, d dice.Dice) (keepers dice.Dice){
	// keep all the dice
	return d
}

func (s FirstRoll) Name() string{
	return "Keep the first roll"
}

