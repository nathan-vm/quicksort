package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

func quicksort(arr []int) []int {
	start := time.Now()
	left, right := 0, len(arr)-1

	stack := [][]int{{left, right}}

	for len(stack) > 0 {
		partition := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		pivotIndex := partition[0]
		pivotValue := arr[pivotIndex]

		i := partition[0]
		for j := partition[0] + 1; j <= partition[1]; j++ {
			if arr[j] < pivotValue {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		arr[i], arr[pivotIndex] = arr[pivotIndex], arr[i]

		if partition[0] < i-1 {
			stack = append(stack, []int{partition[0], i - 1})
		}
		if i+1 < partition[1] {
			stack = append(stack, []int{i + 1, partition[1]})
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
	return arr
}

func main() {
	jsonFile, err := ioutil.ReadFile("input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var arr []int
	json.Unmarshal(jsonFile, &arr)

	arr = quicksort(arr)

}
