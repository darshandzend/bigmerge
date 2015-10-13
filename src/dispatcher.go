package bigmerge

func dispatcher(results chan<- [][]int, input <-chan int) {

	/* if MAXGOROUTINES set
	this means z = current total input size / MAXGOROUTINES per goroutine


	*/

	/*if MAX_N_PER_GOROUTINE set
	this means z = MAX_N_PER_GOROUTINE

	So wait for z inputs to come in
		when z or channel close,
			launch a new sorter goroutine with z inputs, provide results channel


	wait on all goroutines to close

	exit

	*/

}

func sorter(results chan<- [][]int, in []int) {
	/* Choose the sort algorithm for this input size TODO:Make this configurable
		Also if you choose Mergesort here, MERGE-SORTA-CEPTION!
	sort in[]
	put it in results

	exit
	*/

}
