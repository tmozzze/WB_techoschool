package sortgo

import (
	"strconv"
	"testing"

	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_config"
	"github.com/tmozzze/WB_techoschool/L2/L2_10/sortgo_model"
)

// makeLines - make Liens from Strings for tests
func makeLines(data []string) []*sortgo_model.Line {
	var lines []*sortgo_model.Line
	for _, s := range data {
		l := sortgo_model.NewLine(s)
		l.SplitFields("\t") // default \t
		lines = append(lines, l)
	}
	return lines
}

// assertEqual - check equal expected
func assertEqual(t *testing.T, lines []*sortgo_model.Line, expected []string) {
	if len(lines) != len(expected) {
		t.Fatalf("Length mismatch: got %d lines, want %d", len(lines), len(expected))
	}
	for i, line := range lines {
		if line.Raw != expected[i] {
			t.Errorf("index %d: got %q, want %q", i, line.Raw, expected[i])
		}
	}
}

// ------------- FLAG TESTS -------------

func TestCheck(t *testing.T) {
	// Unsorted
	data := []string{
		"banana",
		"apple",
		"cherry",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Check: true}

	if checkSorted(lines, flags) {
		t.Errorf("expected checkSorted to return false for unsorted input, got true")
	}

	// Sorted
	sortedData := []string{
		"apple",
		"banana",
		"cherry",
	}
	lines = makeLines(sortedData)

	if !checkSorted(lines, flags) {
		t.Errorf("expected checkSorted to return true for sorted input, got false")
	}
}

func TestStringSortExpected(t *testing.T) {
	data := []string{
		"banana",
		"apple",
		"cherry",
	}
	expected := []string{
		"apple",
		"banana",
		"cherry",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagKey(t *testing.T) {
	data := []string{
		"1\tbanana",
		"2\tapple",
		"3\tcherry",
	}
	expected := []string{
		"2\tapple",
		"1\tbanana",
		"3\tcherry",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2} // сортировка по колонке 2
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagNum(t *testing.T) {
	data := []string{
		"42",
		"7",
		"100",
	}
	expected := []string{
		"7",
		"42",
		"100",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Num: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagReverse(t *testing.T) {
	data := []string{
		"banana",
		"apple",
		"cherry",
	}
	expected := []string{
		"cherry",
		"banana",
		"apple",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Reverse: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagUnique(t *testing.T) {
	data := []string{
		"apple",
		"banana",
		"apple",
	}
	expected := []string{
		"apple",
		"banana",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Unique: true}
	lines = removeDuplicates(lines, flags)
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagMonth(t *testing.T) {
	data := []string{
		"Jan",
		"Mar",
		"Feb",
		"Dec",
	}
	expected := []string{
		"Jan",
		"Feb",
		"Mar",
		"Dec",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Month: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagIgnoreTrailing(t *testing.T) {
	data := []string{
		"apple  ",
		"banana",
		"  cherry",
		"apple",
	}
	expected := []string{
		"  cherry",
		"apple",
		"apple  ",
		"banana",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{IgnoreTrailing: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestFlagHuman(t *testing.T) {
	data := []string{
		"1K",
		"512",
		"2M",
		"100K",
	}
	expected := []string{
		"512",
		"1K",
		"100K",
		"2M",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Human: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

// ------------- COMBINATION FLAG TESTS -------------

func TestNumReverse(t *testing.T) {
	data := []string{
		"42",
		"7",
		"100",
	}
	expected := []string{
		"100",
		"42",
		"7",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Num: true, Reverse: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyNum(t *testing.T) {
	data := []string{
		"1\t42",
		"2\t7",
		"3\t100",
	}
	expected := []string{
		"2\t7",
		"1\t42",
		"3\t100",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, Num: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyReverse(t *testing.T) {
	data := []string{
		"a\tbanana",
		"b\tapple",
		"c\tcherry",
	}
	expected := []string{
		"c\tcherry",
		"a\tbanana",
		"b\tapple",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, Reverse: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyUnique(t *testing.T) {
	data := []string{
		"1\tapple",
		"2\tbanana",
		"3\tapple",
	}
	expected := []string{
		"1\tapple",
		"2\tbanana",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, Unique: true}
	lines = removeDuplicates(lines, flags)
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyNumReverse(t *testing.T) {
	data := []string{
		"r1\t42",
		"r2\t7",
		"r3\t100",
	}
	expected := []string{
		"r3\t100",
		"r1\t42",
		"r2\t7",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, Num: true, Reverse: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyHumanUnique(t *testing.T) {
	data := []string{
		"a\t1K",
		"b\t512",
		"c\t2M",
		"d\t1K",
	}
	expected := []string{
		"b\t512",
		"a\t1K",
		"c\t2M",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, Human: true, Unique: true}
	lines = removeDuplicates(lines, flags)
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyMonthReverse(t *testing.T) {
	data := []string{
		"row1\tJan",
		"row2\tMar",
		"row3\tFeb",
		"row4\tDec",
	}
	expected := []string{
		"row4\tDec",
		"row2\tMar",
		"row3\tFeb",
		"row1\tJan",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, Month: true, Reverse: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestKeyIgnoreTrailing(t *testing.T) {
	data := []string{
		"a\tapple  ",
		"b\tbanana",
		"c\tapple",
	}
	expected := []string{
		"a\tapple  ",
		"c\tapple",
		"b\tbanana",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Key: 2, IgnoreTrailing: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestUniqueIgnoreTrailing(t *testing.T) {
	data := []string{
		"apple  ",
		"banana",
		"apple",
	}
	expected := []string{
		"apple  ",
		"banana",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Unique: true, IgnoreTrailing: true}
	lines = removeDuplicates(lines, flags)
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestMonthReverse(t *testing.T) {
	data := []string{
		"Jan",
		"Mar",
		"Feb",
		"Dec",
	}
	expected := []string{
		"Dec",
		"Mar",
		"Feb",
		"Jan",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Month: true, Reverse: true}
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

func TestHumanReverseUnique(t *testing.T) {
	data := []string{
		"1K",
		"512",
		"2M",
		"100K",
		"512",
	}
	expected := []string{
		"2M",
		"100K",
		"1K",
		"512",
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{Human: true, Reverse: true, Unique: true}
	lines = removeDuplicates(lines, flags)
	sortLines(lines, flags)

	assertEqual(t, lines, expected)
}

// ------------- PERFOMANCE TESTS -------------

func TestLargeSort(t *testing.T) {
	const n = 100_000 // strings amount

	data := make([]string, n)
	for i := 0; i < n; i++ {
		data[i] = strconv.Itoa(n - i)
	}

	lines := makeLines(data)
	flags := &sortgo_config.Config{
		Num: true,
	}

	t.Logf("Sorting %d lines...", n)

	sortLines(lines, flags)

	if lines[0].Raw != "1" || lines[len(lines)-1].Raw != strconv.Itoa(n) {
		t.Errorf("Large sort failed: first=%s, last=%s", lines[0].Raw, lines[len(lines)-1].Raw)
	}

	t.Logf("First 5 lines: %v", lines[:5])
	t.Logf("Last 5 lines: %v", lines[len(lines)-5:])
}
