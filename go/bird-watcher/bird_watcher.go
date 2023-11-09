package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, element := range birdsPerDay {
		sum += element
	}
	return sum
	//panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	begin := (week - 1) * 7
	end := begin + 7
	
	for i := begin; i < end; i++ {
		sum += birdsPerDay[i]
	}
	return sum
	//panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}
	return birdsPerDay
	//panic("Please implement the FixBirdCountLog() function")
}
