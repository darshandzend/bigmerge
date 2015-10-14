package bigmerge

import "fmt"

func merger(curOp chan<- []int,
	done chan<- struct{},
	results chan []int) {

	for r := range results {
		fmt.Println(r)
	}
	done <- struct{}{}
	curOp <- []int{1, 2, 3}

}

func merge(results chan<- []int, one, two []int) {
	fmt.Println("Merging", one, two)
	results <- append(one, two...)

	/*
		merged := make([]int, len(left)+len(right))

		j, k := 0, 0
		for i := 0; i < len(merged); i++ {
			if j == len(left) {
				for _, d := range right[k:] {
					merged[i] = d
					i++
				}
				break
			}
			if k == len(right) {
				for _, d := range left[j:] {
					merged[i] = d
					i++
				}
				break
			}

			if left[j] < right[k] {
				merged[i] = left[j]
				j++
			} else {
				merged[i] = right[k]
				k++
			}
		}
		results <- merged
	*/

}
