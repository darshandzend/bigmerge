package bigmerge

import "fmt"

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

func (b bigMerge) run() {

	results := make(chan []int, 100)
	go dispatcher(results, b.ip)

	op := make(chan []int)
	go merger(op, b.done, results)

	fmt.Println(<-op)

	//Meahwhile, also wait for user commands

}
