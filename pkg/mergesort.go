package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

func create_array(len int) []int {
	arr := make([]int, len)
	tim := rand.NewSource(time.Now().UnixNano())
	for i := range arr {
		arr[i] = rand.New(tim).Intn(1000)
	}
	return arr
}

func merge(arr *[][2]int, l, m, r int) {

	l_arr := make([][2]int, m-l+1)
	copy(l_arr, (*arr)[l:m+1])

	r_arr := make([][2]int, r-m)
	copy(r_arr, (*arr)[m+1:r+1])

	i, j, k := 0, 0, l
	for i < len(r_arr) && j < len(l_arr) {
		if r_arr[i][1] < l_arr[j][1] {
			(*arr)[k] = r_arr[i]
			i++
		} else {
			(*arr)[k] = l_arr[j]
			j++
		}
		k++
	}

	for i < len(r_arr) {
		(*arr)[k] = r_arr[i]
		i++
		k++
	}

	for j < len(l_arr) {
		(*arr)[k] = l_arr[j]
		j++
		k++
	}

}

func (s *Sorter) mergs(tupl *[][2]int, l, r int) {
	m := (l + r) / 2
	if l < r {
		s.mergs(tupl, l, m)
		s.mergs(tupl, m+1, r)

		var aux Merger
		aux.Merge = [2][]int{s.GetIndexes((*tupl)[l : m+1]), s.GetIndexes((*tupl)[m+1 : r+1])}
		merge(tupl, l, m, r)
		aux.Into = s.GetIndexes((*tupl)[l : r+1])
		s.MergeList = append(s.MergeList, aux)
	}
}

func createTuple(arr []int) [][2]int {
	tupleArr := make([][2]int, len(arr))
	for idx, val := range arr {
		tupleArr[idx] = [2]int{idx, val}
	}
	return tupleArr
}

func (s *Sorter) GetIndexes(tupl [][2]int) []int {
	arr := make([]int, len(tupl))
	for idx, val := range tupl {
		arr[idx] = val[0]
	}
	return arr

}
func (s *Sorter) GetValues(tupl [][2]int) []int {
	arr := make([]int, len(tupl))
	for idx, val := range tupl {
		arr[idx] = val[1]
	}
	return arr
}

func (s *Sorter) Sort(arr []int) [][2]int {
	s.tuple = createTuple(arr)
	s.mergs(&s.tuple, 0, len(arr)-1)
	s.formatResult(s.tuple)
	fmt.Println(s.MergeList)
	return s.tuple
}

func (s *Sorter) formatResult(arr [][2]int) []Slot {
	s.SlotList = make([]Slot, len(arr))
	for idx, val := range arr {
		s.SlotList[idx].Idx = val[0]
		s.SlotList[idx].Value = val[1]
	}
	fmt.Println(s.SlotList)
	return s.SlotList
}
