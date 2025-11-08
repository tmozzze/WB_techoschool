package comporators

import (
	"strconv"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"
)

func NumericComporator(key int) Comporator {
	return func(i, j *model.Line) bool {
		keyI := i.GetField(key)
		keyJ := j.GetField(key)

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

		return numI < numJ
	}
}
