package dice

import (
	"math"
	"sort"
	//"fmt"
	)

type Dice []Die

func NewDice(num int) Dice{
	return Dice(make([]Die, num))
}

func FromDie(die Die) Dice{
	dice := NewDice(1)
	dice[0] = die
	return dice
}

func (d Dice) Roll() {
	
	for i := 0; i < len(d); i++{
		d[i] = RollDie()
	}
	// sort dice
	sort.Sort(d)
	
}

func (d Dice) PointValue() int{

	if len(d)==6{
		switch {
		
			// all single number
			case d.IsSingleNumber():
				return math.MaxInt32
				
			// 123456
			case d.IsRun():
				return 1000
			
		}
	}
	
	points := 0
	
	// three of kind
	triple1, triple2, leftovers := d.Triples()
	if len(triple1)>0{
		points += triple1[0].TriplePointValue()
	}
	if len(triple2)>0{
		points += triple2[0].TriplePointValue()
	}
	
	// single scores	
	for _,single := range leftovers{
		points += single.PointValue()
	}
	
	return points;
}

func(d Dice) Nonscoring() Dice{
	switch{
		case d.IsSingleNumber() && len(d) == 6:
			return NewDice(0)
		case d.IsRun() || (d.IsSingleNumber()):
			return NewDice(0)
	}
	_, _, leftovers := d.Triples()
	return leftovers.Subtract([]Die{Die(5),Die(1)})
}

func (d Dice) Triples() (triple1 Dice, triple2 Dice, leftovers Dice){

	//triple1, triple2 := d.pips[0:0], d.pips[0:0]
	//lo := make([]int,6)

	repeats := 0
	previousVal := Die(math.MaxUint8) 
	for i,pips := range d {
		
		switch{
			case pips == previousVal:
				repeats++
			case pips != previousVal:
				previousVal = pips
				for _,val := range d[i-repeats:i]{
					leftovers = append(leftovers,val) 
				}
				repeats = 1
		}
		
		if repeats == 3{
			switch{
				case len(triple1) == 0:
					triple1 = d[i-2:i+1]
					repeats = 0
				case len(triple2) == 0:
					triple2 = d[i-2:i+1]
					repeats = 0
			}
		}
		
		if i==(len(d)-1) && repeats > 0{
			// this is the last die and it was not part of a triple--add it in
			for _,val := range d[i-repeats+1:i+1]{
				leftovers = append(leftovers,val) 
			}
		}
		
	} // end for
	
	return triple1, triple2, leftovers
}

func (d Dice) IsSingleNumber() bool{
	if len(d) == 0{
		return true
	}
	
	singleValue := d[0]
	for _,pips := range d{
		if pips != singleValue{
			return false
		}
	}
	return true
}

func (d Dice) IsRun() bool{
	return  len(d) == 6 &&
			d[0] == 1 && 
			d[1] == 2 && 
			d[2] == 3 && 
			d[3] == 4 && 
			d[4] == 5 && 
			d[5] == 6
}

func (d Dice) Subtract(q Dice) Dice{
	keepers := make([]Die, 0)
	// return all die that are in d but NOT q
	sort.Sort(q)
	for _,dv := range d{
		qi := sort.Search(len(q), func(i int) bool{return q[i]>=dv})
		if qi == len(q) || q[qi]!=dv {
			// this d element was not found in q
			keepers = append(keepers, dv)
		}
	}
	diff := NewDice(len(keepers))
	for i,v := range keepers{
		diff[i] = v
	} 
	//fmt.Printf("\n\np: %v",p)
	//fmt.Printf("\nq: %v",q)
	//fmt.Printf("\ndiff: %v",diff)
	return diff
}

// Methods necessary for sort.Inteface
func (d Dice) Len() int{
	return len(d)
}
func (d Dice) Less(i, j int) bool{
	return d[i] < d[j]
}
func (d Dice) Swap(i, j int){
	d[i], d[j] = d[j], d[i]
}

