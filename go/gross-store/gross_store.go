package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	result := make(map[string]int)

	result["quarter_of_a_dozen"] = 3
	result["half_of_a_dozen"] = 6
	result["dozen"] = 12
	result["small_gross"] = 120
	result["gross"] = 144
	result["great_gross"] = 1728

	return result
	//panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
	//panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	_, billExists := bill[item]
	if billExists {
		bill[item] += unitValue
	} else {
		bill[item] = unitValue
	}
	
	return true
	//panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemValue, itemExists := bill[item]
	unitsValue, unitsExists := units[unit]
	
	if !itemExists || ! unitsExists {
		return false
	}
	
	if itemValue - unitsValue < 0 {
		return false
	}
	
	if itemValue - unitsValue == 0 {
		delete(bill, item)
	} else {
		bill[item] -= units[unit]
	}
	
	return true
	//panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	
	value, exists := bill[item]
	
	if !exists {
		return 0, false
	}
	return value, true
	//panic("Please implement the GetItem() function")
}
