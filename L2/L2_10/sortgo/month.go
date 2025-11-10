package sortgo

import (
	"fmt"
	"strings"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
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

func monthSort(lineI, lineJ *model.Line, flags *config.Config) bool {
	keyI := getSortKey(lineI, flags) // When -k given --> field
	keyJ := getSortKey(lineJ, flags) // When -k default --> raw

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

	if monthI == monthJ {
		return lineI.Raw < lineJ.Raw
	}
	// When allright
	return monthI < monthJ
}
