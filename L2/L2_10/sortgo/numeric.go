package sortgo

import (
	"strconv"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

func numericSort(lineI, lineJ *model.Line, flags *config.Config) bool {
	keyI := getSortKey(lineI, flags) // When -k given --> field
	keyJ := getSortKey(lineJ, flags) // When -k default --> raw

	numI, errI := strconv.ParseFloat(keyI, 64)
	numJ, errJ := strconv.ParseFloat(keyJ, 64)

	// When 1st and 2nd are not num
	if errI != nil && errJ != nil {
		return keyI < keyJ
	}
	// When 1st not num
	if errI != nil {
		return true
	}
	// When 2nd not num
	if errJ != nil {
		return false
	}

	if numI == numJ {
		return lineI.Raw < lineJ.Raw
	}
	return numI < numJ
}
