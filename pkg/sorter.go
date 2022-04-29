package pkg

type Slot struct {
	Idx   int
	Value int
}

type Merger struct {
	Merge [2][]int
	Into  []int
}

type iSorter interface {
	Sort(arr []int) [][2]int
	GetIndexes(tupl [][2]int) []int
	GetValues(tupl [][2]int) []int
	formatResult(arr [][2]int) []Slot
}

type Sorter struct {
	Arr       []int
	SlotList  []Slot
	tuple     [][2]int
	MergeList []Merger
}
