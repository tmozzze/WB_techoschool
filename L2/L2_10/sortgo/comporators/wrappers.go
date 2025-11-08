package comporators

import "github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"

// ReverseComparator - invert
func ReverseComporator(base Comporator) Comporator {
	return func(i, j *model.Line) bool {
		return !base(i, j)
	}
}

func StableTieComporator(base Comporator) Comporator {
	return func(i, j *model.Line) bool {
		if !base(i, j) && !base(j, i) {
			return i.Raw < j.Raw
		}
		return base(i, j)
	}
}
