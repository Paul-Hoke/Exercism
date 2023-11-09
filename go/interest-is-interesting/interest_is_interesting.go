package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return .5
	case balance >= 1000 && balance < 5000:
		return 1.621
	default:
		return 2.475
	}
	
	//panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	switch {
	case balance < 0:
		return balance * (float64(InterestRate(balance)) / 100)
	case balance >= 0 && balance < 1000:
		return balance * (float64(InterestRate(balance)) / 100)
	case balance >= 1000 && balance < 5000:
		return balance * (float64(InterestRate(balance)) / 100)
	default:
		return balance * (float64(InterestRate(balance)) / 100)
	}
	//panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	balance += Interest(balance)
	return balance
	//panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	yearsBeforeDesiredBalance := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		yearsBeforeDesiredBalance++
	}
	return yearsBeforeDesiredBalance
	//panic("Please implement the YearsBeforeDesiredBalance function")
}
