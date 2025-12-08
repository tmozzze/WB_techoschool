package cutgo_model

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type IntDiapozoneValue []int

func (s *IntDiapozoneValue) Type() string {
	return "intDiapozone"
}

func (s *IntDiapozoneValue) String() string {
	var strSlice []string
	for _, v := range *s {
		strSlice = append(strSlice, strconv.Itoa(v))
	}
	return strings.Join(strSlice, ",")
}

func (s *IntDiapozoneValue) Set(value string) error {
	ranges := strings.Split(value, ",")
	for _, r := range ranges {
		if strings.Contains(r, "-") {
			parts := strings.Split(r, "-")
			if len(parts) != 2 {
				return fmt.Errorf("invalid diapozone format: %s", r)
			}
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return fmt.Errorf("invalid start format: %s", parts[0])
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("invalid end format: %s", parts[1])
			}
			if start > end {
				return fmt.Errorf("invalid diapozone: start %d cannot be greater then end %d", start, end)
			}
			for i := start; i <= end; i++ {
				*s = append(*s, i)
			}

		} else {
			val, err := strconv.Atoi(r)
			if err != nil {
				return fmt.Errorf("invalid value: %s", r)
			}
			*s = append(*s, val)
		}
	}
	return nil
}

// Len - return len of slice
func (s *IntDiapozoneValue) Len() int {
	return len(*s)
}

// GetSlice - return slice
func (s *IntDiapozoneValue) GetSlice() []int {
	return *s
}

// SortFields - sorting slice of fields if NOT sorted
func (s *IntDiapozoneValue) SortFields() {
	sort.Ints(*s)

	s.removeDuplicates()
}

func (s *IntDiapozoneValue) removeDuplicates() {
	if len(*s) < 2 {
		return
	}

	uniqueIdx := 0
	for i := 1; i < len(*s); i++ {
		if (*s)[uniqueIdx] != (*s)[i] {
			uniqueIdx++
			(*s)[uniqueIdx] = (*s)[i]
		}
	}
	*s = (*s)[:uniqueIdx+1]

}
