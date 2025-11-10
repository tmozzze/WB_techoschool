package sortgo

import (
	"github.com/tmozzze/WB_techoschool/L2/L2_10/config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/model"
)

func stringSort(lineI, lineJ *model.Line, flags *config.Config) bool {
	keyI := getSortKey(lineI, flags) // When -k given --> field
	keyJ := getSortKey(lineJ, flags) // When -k default --> raw

	if keyI == keyJ {
		return lineI.Raw < lineJ.Raw
	}

	return keyI < keyJ
}
