package main

import (
	"farkle"
	"fmt"
	"math"
	"farkle/strategies"
	)

const(
	DEBUG = false
)

func main(){
	// Setup
	games := 10000
	fmt.Printf("%v games of farkle",games)

	// Strategies
	strats := []farkle.Strategy{strategies.NewFirstRoll()}

	for i:=200; i<500; i+=50{
		strats = append(strats,strategies.NewStopAtN(i))
	}

	// Eval
	for _,strategy := range strats{
		fmt.Printf("\n\n\"%v\"",strategy.Name())
		fmt.Printf("\nAverage Score: %v",averageStrategyScore(strategy,games))
	}
}

func averageStrategyScore(s farkle.Strategy, games int) float64{

	cumulativeScore := 0
	allSameNumberWins := 0

	for i := 0; i<games; i++{
		
		if DEBUG{ fmt.Printf("\n\nGame %v",i) }
		
		play := farkle.GameFactory(s)
		for gameScore, finished := play(); true; gameScore, finished = play() {

			// jackpot!
			if gameScore==math.MaxInt32{
				allSameNumberWins++
			}
			
			if finished{
				cumulativeScore += gameScore
				if DEBUG{ fmt.Printf("\nfinal score: %v",gameScore) }
				break
			}
			
		}
		
		
		
	}
	
	return float64(cumulativeScore) / float64(games-allSameNumberWins)
	
}