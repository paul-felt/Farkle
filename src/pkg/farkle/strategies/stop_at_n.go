package strategies

import (
	"dice"
	"fmt"
)

/* 
 * Play until you get at least N Points, then stop 
 */
type StopAtN struct{
	n int
}

func NewStopAtN(n int) *StopAtN{
	return &StopAtN{n}
}

func (s *StopAtN) Keep(score int, d dice.Dice) (keepers dice.Dice){
	// keep all 
	switch{
		// keep jackpot
		case d.IsSingleNumber() && len(d)==6:
			return d
		// keep run
		case d.IsRun():
			return d
		// all the dice are used
		case len(d.Nonscoring())==0:
			return d
		// score is higher than our threshold
		case score+d.PointValue() >= s.n:
			return d
	}
	
	// keep some 
	
	// if there's a triple over 200, keep that
	triple1,_,_ := d.Triples()
	if triple1 != nil && triple1.PointValue()>=300{
		return triple1
	}
	
	// otherwise, keep a single die (max(1,5))
	singleIndex := bestSingleIndex(d)
	if (singleIndex!=-1){
		return dice.FromDie(d[singleIndex])
	} 
	
	// otherwise, keep the 200 triple
	return triple1

	//// keep all scoring dice
	//return d.Subtract(d.Nonscoring())
	
}

func (s *StopAtN) Name() string{
	return fmt.Sprintf("StopAt%v",s.n)
}

func bestSingleIndex(d dice.Dice) int{
	five := -1
	for i,v := range d{
		switch v {
			case 1:
				return i;
			case 5:
				five = i 
		}
	}
	return five
}
