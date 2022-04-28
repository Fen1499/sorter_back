package pkg

type Sorter struct {
	ArrMap map[int]int
}

func (s *Sorter) InitMap(arr []int) map[int]int {
	s.ArrMap = make(map[int]int)
	for idx, key := range arr {
		s.ArrMap[idx] = key
	}
	return s.ArrMap
}

func (s Sorter) printIndexes(arr []int) []int {
	printArr := make([]int, len(arr))
	for idx, key := range arr {
		printArr[idx] = s.ArrMap[key]
	}
	//fmt.Println(arr)
	return printArr
}
