package sortgo

import (
	"fmt"
	"strings"
)

func getMonthValue(monthName string) (int, error) {
	var monthOrder = map[string]int{
		"jan": 1, "feb": 2, "mar": 3, "apr": 4,
		"may": 5, "jun": 6, "jul": 7, "aug": 8,
		"sep": 9, "oct": 10, "nov": 11, "dec": 12,
	}
	// Check Lenght
	if len(monthName) < 3 {
		return 0, fmt.Errorf("incorrect string. Lenght string must be 3 or higher")
	}
	// Formating month name
	formatedMonthName := strings.ToLower(monthName[:3])

	// Get value by name of month
	value := monthOrder[formatedMonthName]
	if value == 0 {
		return 0, fmt.Errorf("incorrect string")
	}
	return value, nil
}

func monthSort(keyI, keyJ string) bool {
	monthI, errI := getMonthValue(keyI)
	monthJ, errJ := getMonthValue(keyJ)

	// When 1st and 2nd are not month
	if errI != nil && errJ != nil {
		return keyI < keyJ
	}
	// When 1st are not month
	if errI != nil {
		return true
	}
	// When 2nd are not month
	if errJ != nil {
		return false
	}
	// When allright
	return monthI < monthJ
}
