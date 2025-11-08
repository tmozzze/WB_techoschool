package comporators

import (
	"fmt"
	"strings"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"
)

// monthMap - int value of month's name
var monthMap = map[string]int{
	"jan": 1, "feb": 2, "mar": 3, "apr": 4,
	"may": 5, "jun": 6, "jul": 7, "aug": 8,
	"sep": 9, "oct": 10, "nov": 11, "dec": 12,
}

func MonthComporator(key int) Comporator {
	return func(i, j *model.Line) bool {
		keyI := i.GetField(key)
		keyJ := j.GetField(key)

		monthI, errI := parseMonth(keyI)
		monthJ, errJ := parseMonth(keyJ)

		// When 1st and 2nd are not month
		if errI != nil && errJ != nil {
			return keyI < keyJ
		}
		// When 1st are not month
		if errI != nil {
			return true
		}
		// When 1st are not month
		if errJ != nil {
			return false
		}
		// When allright
		return monthI < monthJ
	}
}

func parseMonth(monthName string) (int, error) {
	// Check Lenght
	if len(monthName) < 3 {
		return 0, fmt.Errorf("incorrect string. Lenght string must be 3 or higher")
	}
	// Formating month name
	formatedMonthName := strings.ToLower(monthName[:3])

	// Get value by name of month
	value := monthMap[formatedMonthName]
	if value == 0 {
		return 0, fmt.Errorf("incorrect string")
	}
	return value, nil
}
