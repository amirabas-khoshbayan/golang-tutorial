package main

import (
	"fmt"
	"sync"
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

	blue := worker{"blue"}
	purple := worker{"purple"}
	orange := worker{"orange"}

	var wg sync.WaitGroup
	wg.Add(3)

	go blue.doJobs(jobs[:2], &wg)
	go purple.doJobs(jobs[2:4], &wg)
	go orange.doJobs(jobs[4:6], &wg)
	wg.Wait()

}

func (w worker) doJobs(jobs []string, wg *sync.WaitGroup) {
	for _, job := range jobs {
		time.Sleep(time.Second * 1)
		fmt.Println(job, "by", w.color, "is done")
	}
	wg.Done()
}
