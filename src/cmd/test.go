package main

import (
	"fmt"
	"dice"
)

func main(){
	d := dice.Dice([]dice.Die{2,2,2,5,5,5})
	
	triple1, triple2, leftovers := d.Triples()
	
	fmt.Printf("t1: %v\tt2: %v\tt3: %v", triple1, triple2, leftovers)
	
}