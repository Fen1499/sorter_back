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

func merge(arr *[]int, l, m, r int) {

	l_arr := make([]int, m-l+1)
	copy(l_arr, (*arr)[l:m+1])

	r_arr := make([]int, r-m)
	copy(r_arr, (*arr)[m+1:r+1])

	i, j, k := 0, 0, l
	for i < len(r_arr) && j < len(l_arr) {
		if r_arr[i] < l_arr[j] {
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

func Mergs(arr *[]int, l, r int) {
	m := (l + r) / 2
	if l < r {
		Mergs(arr, l, m)
		Mergs(arr, m+1, r)
		fmt.Println("merge: ", (*arr)[l:m+1], (*arr)[m+1:r+1])
		merge(arr, l, m, r)
		fmt.Println("into: ", (*arr)[l:r+1])
	}
}

func createTuple(arr []int) [][2]int {
	tupleArr := make([][2]int, len(arr))
	for idx, val := range arr {
		tupleArr[idx] = [2]int{idx, val}
	}
	fmt.Println(tupleArr)
	return tupleArr
}

// func Sort(arr []int) {
// 	tupl := createTuple(arr)
// 	Mergs(arr, 0, len(arr)-1)
// }
