package comporators

import "github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo/model"

func StringComporator(key int) Comporator {
	return func(i, j *model.Line) bool {
		return i.GetField(key) < j.GetField(key)
	}
}
