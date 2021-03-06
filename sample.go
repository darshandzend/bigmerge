package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/darshandzend/bigmerge/src"
)

func main() {
	b := bigmerge.New()
	ch := b.IpChan()
	for i := 100; i > 0; i-- {
		ch <- rand.Intn(100)
		time.Sleep(time.Millisecond * 1)
	}
	close(ch)
	<-b.Done()
	fmt.Println("Done.")
}
