package bigmerge

import "log"

type BigMerger interface {
	IpChan() chan<- int
	CurOp() []int
	Done() <-chan struct{}
	//Stop
	//RemoveDups
}

type bigMerge struct {
	ip    chan int
	curOp []int
	done  chan struct{}
}

func New() BigMerger {
	ip := make(chan int)
	op := make([]int, 0)
	done := make(chan struct{})
	b := bigMerge{ip, op, done}
	go b.run()
	return b
}

func (b bigMerge) IpChan() chan<- int    { return b.ip }
func (b bigMerge) CurOp() []int          { return b.curOp }
func (b bigMerge) Done() <-chan struct{} { return b.done }

type auxResult struct {
	auxr  chan []int
	count int
}

func (a *auxResult) insert(arr []int) {
	a.count++
	a.auxr <- arr
}

func (a *auxResult) remove() []int {
	a.count--
	return <-a.auxr
}

func (b bigMerge) run() {

	results := make(chan []int, 100)
	go dispatcher(results, b.ip)

	sorted := make(chan []int)
	go merger(sorted, results)

	log.Println(<-sorted)
	b.done <- struct{}{}

	//Meahwhile, also wait for user commands

}
