package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate * .01)
	//panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
	//panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groups := uint(carsCount) / 10
	groupsCost := groups * 95000
	
	individualCars := uint(carsCount) - (groups * 10)
	individualCarsCost := individualCars * 10000
	
	return groupsCost + individualCarsCost
	//panic("CalculateCost not implemented")
}


//CalculateWorkingCarsPerHour(221, 100.000000) = 22100.000000, want 221.000000