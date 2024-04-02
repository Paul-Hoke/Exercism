package expenses

import (
	"errors"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	
	result := []Record{}
	
	for _, v := range in {
		if(predicate(v)) {
			result = append(result, v)
		}
	}
	
	return result
	//panic("Please implement the Filter function")
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	
	return func(r Record) bool {
		if(r.Day >= p.From && r.Day <= p.To) {
			return true
		}
		return false
	}
	//panic("Please implement the ByDaysPeriod function")
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	
	return func(r Record) bool {
		return c == r.Category
	}
	//panic("Please implement the ByCategory function")
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	
	result := 0.0
	filteredSlice := Filter(in, ByDaysPeriod(p))
	
	for _, v := range filteredSlice {
		result += v.Amount
	}
	
	return result
	//panic("Please implement the TotalByPeriod function")
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	
	if (len(Filter(in, ByCategory(c))) <= 0) {
		return 0.0, errors.New("No Records existed for the specified catgegory")
	}
	
	return TotalByPeriod(Filter(Filter(in, ByCategory(c)), ByDaysPeriod(p)), p), nil
	
	//panic("Please implement the CategoryExpenses function")
}
