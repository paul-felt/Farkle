package farkle

import(
	"dice"
	"fmt"
	)


const(
	DEBUG = false
)

type Game struct{
	score int
	dice dice.Dice
}

type Strategy interface{
	Keep(score int, d dice.Dice) (keepers dice.Dice)
	Name() string
}

func GameFactory(s Strategy) func() (cumulativeScore int, finished bool){
		
	game := Game{0, dice.NewDice(6)}
	// play a single round using strategy s and return resulting score
	return func() (score int, finished bool){
	
		// roll dice
		if DEBUG{ fmt.Printf("\n\trolling %v dice . . .", len(game.dice)) }
		game.dice.Roll()
		if DEBUG{ fmt.Print(game.dice) }
		
		// farkle!
		if game.dice.PointValue() == 0{
			if DEBUG{ fmt.Printf("\n\tfarkle! . . . %v", game.dice.PointValue()) }
			switch{
				// first roll you get pity points 
				case len(game.dice) == 6:
					return 50, true
				case len(game.dice) < 6:
					return 0, true
			}
		}
		
		// strategy decides which to keep
		keepers := s.Keep(game.score, game.dice)
		game.dice = game.dice.Subtract(keepers)
		
		// If they successfully used all dice, they get another 6 dice
		if len(game.dice)==0 && len(keepers.Nonscoring())==0{
			game.dice = dice.NewDice(6)
			if DEBUG{ fmt.Print("\tCongratulations! You used all your dice. Have another 6.") }
		}
		
		// score for next round
		game.score += keepers.PointValue()
		if DEBUG{ fmt.Printf("\n\tkeep %v for %v points (%v)", keepers, keepers.PointValue(), game.score) }
		if DEBUG{ fmt.Printf("\n\troll again %v",game.dice) }
	
		return game.score, len(game.dice) == 0
	}
	
}

