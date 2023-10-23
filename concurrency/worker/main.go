package main

import (
	"fmt"
	"time"
)

type worker struct {
	color string
}

// A goroutine is a lightweight thread managed by the Go runtime (created in µs, around 2.6 KB memory)
// Go runtime is a part of Go executable that manages our Go program (e.g. memory allocation, channels communication, goroutines creation) and Communicate with OS kernel
// A Thread is a sequence of independent instructions needs to be processed by CPU

func main() {
	jobs := []string{"Job 1", "Job 2", "Job 3", "Job 4", "Job 5", "Job 6"}
	blue, purple, orange := worker{"blue"}, worker{"purple"}, worker{"orange"}

	doneJobs := make(chan string)

	go blue.doJobs(jobs[:2], doneJobs)
	go purple.doJobs(jobs[2:4], doneJobs)
	go orange.doJobs(jobs[4:6], doneJobs)

	for range jobs {
		fmt.Println(<-doneJobs)
	}
	//for  job := range doneJobs {
	//	fmt.Println(job)
	//}
}

func (w worker) doJobs(jobs []string, doneJobs chan string) {
	for _, job := range jobs {
		time.Sleep(time.Second * 1)
		doneJobs <- fmt.Sprint(job, " by ", w.color, " is done")
	}
	//close(doneJobs)

}
