package bigmerge

import (
	"fmt"
	"sort"
	"sync"
)

const MAXPERGO = 10

func dispatcher(results chan<- []int, input <-chan int) {

	nextBatch := make([]int, MAXPERGO)
	var sorters sync.WaitGroup //TODO: Is this necessary?
	i := 0
	for n := range input {
		nextBatch[i] = n
		i++
		if i == MAXPERGO {
			sorters.Add(1)
			go sorter(results, nextBatch, &sorters)
			i = 0
			nextBatch = make([]int, MAXPERGO)
		}
	}
	//process the rest
	if i != 0 {
		sorters.Add(1)
		go sorter(results, nextBatch, &sorters)
	}

	sorters.Wait()
	fmt.Println("All inputs dispatched")
	close(results)
}

func sorter(results chan<- []int, in []int, sorters *sync.WaitGroup) {
	defer sorters.Done()
	sort.Ints(in)
	results <- in
}
