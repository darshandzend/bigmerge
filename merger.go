package bigmerge

func merger(results chan [][]int) {
	/*
		The master merger.

		wait on results for any new slice to come in

		EITHER
		decide on how many mergers to launch
			nom = n / threshold no. of slices to merge
		launch nom mergers, give them threshold arrays and results channel
		OR
		take first two slices, merge, put it back in results.
			(also update currentResult)

		CHANNEL LOOP! CAREFUL WITH DEADLOCKS!!

		when no more slices to merge, exit.
	*/
}

func merge(one, two []int) []int {
	/*
		simple merge function
	*/
}
