package pkg

type Slot struct {
	Idx   int
	Value int
}

type ISorter interface {
	Sort(arr []int) [][2]int
	GetIndexes(tupl [][2]int) []int
	GetValues(tupl [][2]int) []int
	GetResult() [][]int
}

type Sorter struct {
	Arr      []int
	SlotList []Slot
	tuple    [][2]int
	Tree     [][]int
}
