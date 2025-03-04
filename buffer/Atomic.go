package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
main

import (
"fmt"
"sync"
"sync/atomic"
)

func main() {

	var count atomic.Uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				count.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count.Load())

}

