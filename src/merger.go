package bigmerge

import (
	"log"
	"sync"
	"time"
)

func merger(sorted chan<- []int, parts chan []int) {

	done := make(chan bool)
	var res []int
	var wg sync.WaitGroup

	go func() {
		var lastResult []int
		for {
			sl1, ok := <-parts
			if !ok {
				break
			}
			lastResult = sl1
			sl2, ok := <-parts
			if !ok {
				lastResult = merge(sl1, sl2)
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				parts <- merge(sl1, sl2)
			}()
		}
		log.Println("Shutdown")
		res = lastResult
		done <- true
	}()

	//TODO: Find a non-hacky solution to stop wait() from returning prematurely
	time.Sleep(time.Second)

	go func() {
		wg.Wait()
		log.Println("All done.")
		close(parts)
	}()
	<-done
	sorted <- res
}

func merge(left, right []int) []int {
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

	return merged
}

/*

func merger(sorted chan<- []int, results chan []int) {
	//sortedSnapshot := make([]int, 0) //TODO:size?

	auxResults := make(chan []int, 100)

	go func() {
		for r := range results { //to keep the close() cascade cnsistant
			auxResults <- r
		}
	}()
	//var mergers sync.WaitGroup

	var ar1 []int

	for {
		//temporary hack to make things work for now
		ar1 = <-auxResults
		log.Println("ar1---", ar1)
		if auxEmpty(auxResults) {
			log.Println("ar1", ar1)
			sorted <- ar1
		}

		ar2 := <-auxResults
		go merge(auxResults, ar1, ar2)
	}

	/*	go func() {
			ar1, more := <-auxResults
			mergers.Add(1)
			log.Println("ar1, more", ar1, more)
			if !more {
				//closed
				sorted <- ar1
				return
			}
			ar2, more := <-auxResults
			if !more {
				//closed
				sorted <- append(ar1, ar2...)
				return
			}
			go merge(auxResults, &mergers, ar1, ar2)
		}()

		mergers.Wait()
		//everyone's closed.
		close(auxResults)

		sorted <- sortedSnapshot
}


func merge(results chan<- []int, one, two []int) {
	log.Println("Merging", one, two)
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

}

func auxEmpty(a chan []int) bool {
	if l := len(a); l == 0 {
		time.Sleep(time.Second * 5)
		if l = len(a); l == 0 {
			return false
		}
	}
	return true
}
*/
