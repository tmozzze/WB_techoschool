package sortgo

import (
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_model"
)

func removeDuplicates(lines []*sortgo_model.Line, flags *sortgo_config.Config) []*sortgo_model.Line {
	var unique []*sortgo_model.Line
	uniqueMap := make(map[string]struct{})

	for _, line := range lines {
		key := getSortKey(line, flags)

		if _, exists := uniqueMap[key]; exists {
			continue
		}
		uniqueMap[key] = struct{}{}
		unique = append(unique, line)

	}
	return unique
}
