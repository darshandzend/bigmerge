package bigmerge

type BigMerger interface {
	IpChannel() chan<- int
	CurOp() []int
	//Stop
	//RemoveDups
}

type bigMerge struct {
	ip    chan int
	curOp []int
}

func New() BigMerger {
	ip := make(chan int)
	op := make([]int) //TODO:capacity, size?
	go b.run()
	b := BigMerge{ip, op}
}

func (b bigMerge) IpChan() chan<- int { return b.ip }
func (b bigMerge) CurOp() []int       { return b.curOp }

func (b BigMerge) run() {

	//wait for input in inputs channel, close when no inputs

	//launch dispatcher

	//launch master merger

	//wait for all to complete

	//Meahwhile, also wait for user commands

	//when all results, show output
}
