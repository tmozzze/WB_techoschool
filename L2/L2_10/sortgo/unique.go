package sortgo

import (
	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

func removeDuplicates(lines []*model.Line, flags *config.Config) []*model.Line {
	var unique []*model.Line
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
