package pkg

type Slot struct {
	idx   int
	value int
}

type iSorter interface {
	Sort(arr []int) [][2]int
	GetIndexes(tupl [][2]int) []int
	GetValues(tupl [][2]int) []int
}

type Sorter struct {
	Arr   []int
	Slot  []Slot
	tuple [][2]int
}
