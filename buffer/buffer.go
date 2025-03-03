package main

import "fmt"

func main() {
	// lets create two channels
	jobs := make(chan int, 5)
	done := make(chan bool)

	//reciever routine
	go func() {

		for {
			j, open := <-jobs
			if open {
				fmt.Println("job recieved", j)
			} else {
				fmt.Println("All jobs has been sent")
				done <- true
				return
			}
		}
	}()

	for i := 0; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	close(jobs)
	fmt.Println("Sent all jobs")
	<-done
	_, ok := <-jobs
	fmt.Println("recieved more jobs:", ok)

}
