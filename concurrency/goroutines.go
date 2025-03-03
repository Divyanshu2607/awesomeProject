package m

import (
	"fmt"
	"sync"
	"time"
)

func f(r string) {
	for i := 0; i < 3; i++ {
		fmt.Println("r", ":", i)
	}
}

func worker(ch2 chan bool) {
	fmt.Print("worker is working..")
	time.Sleep(time.Second)
	fmt.Println("work is done")
	ch2 <- true
}
func ft(msg string, doner chan<- string) {
	doner <- msg
}
func transfer(reciever chan<- string, doner <-chan string) {
	msg := <-doner
	reciever <- msg

}
func main() {
	f("direct")
	go f("hello")
	func(name string) {
		fmt.Println(name)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	ch := make(chan string)
	go func() {
		ch <- "Message"

	}()

	msg := <-ch
	fmt.Println(msg)

	ch2 := make(chan bool, 1)
	go worker(ch2)
	<-ch2
	reciever := make(chan string, 1)
	doner := make(chan string, 1)
	ft("hello", doner)
	transfer(reciever, doner)
	fmt.Println(<-reciever)

	c1 := make(chan string)
	c2 := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Message transfered to channel 1"
		close(c1)
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Message transfered to channel 2"
		close(c2)
		wg.Done()
	}()

	//one way is to make an infite loop
	//for {
	//	select {
	//	case msg1, open := <-c1:
	//		if open {
	//			fmt.Println(msg1)
	//		} else {
	//			c1 = nil
	//		}
	//
	//	case msg2, open := <-c2:
	//		if open {
	//			fmt.Println(msg2)
	//		} else {
	//			c2 = nil
	//		}
	//	default:
	//		if c1 == nil && c2 == nil {
	//			fmt.Println("All the messages were captured")
	//			return
	//		}
	//
	//	}
	//}

	//second way is to use sync.WaitGroup
	go func() {
		for i := range c1 {
			fmt.Println(i)
		}
	}()

	go func() {
		for i := range c2 {
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("all message has been printed")
}
