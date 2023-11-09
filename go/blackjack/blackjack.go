package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	
	switch {
	case card =="ace":
		return 11
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		return 10
	default:
		return 0
	}
	
	//panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	
	parseCard1 := ParseCard(card1)
	parseCard2 := ParseCard(card2)
	parseDealerCard := ParseCard(dealerCard)
	
	switch {
	case parseCard1 == 11 && parseCard2 == 11:
		return "P"
	case parseCard1 + parseCard2 == 21 && parseDealerCard < 10:
		return "W"
	case parseCard1 + parseCard2 == 21 && parseDealerCard >= 10:
		return "S"
	case parseCard1 + parseCard2 >= 17 && parseCard1 + parseCard2 <= 20:
		return "S"
	case parseCard1 + parseCard2 >= 12 && parseCard1 + parseCard2 <= 16 && parseDealerCard >= 7:
		return "H"
	case parseCard1 + parseCard2 >= 12 && parseCard1 + parseCard2 <= 16:
		return "S"
	default:
		return "H"
	}
	//panic("Please implement the FirstTurn function")
}
