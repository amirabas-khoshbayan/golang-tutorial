package main

import (
	"fmt"
	"time"
)

type worker struct {
	color string
}

func main() {
	jobs := []string{"Job 1", "Job 2", "Job 3", "Job 4", "Job 5", "Job 6"}

	blue := worker{"blue"}
	purple := worker{"purple"}
	orange := worker{"orange"}

	blue.doJobs(jobs[:2])
	purple.doJobs(jobs[2:4])
	orange.doJobs(jobs[4:6])
}

func (w worker) doJobs(jobs []string) {
	for _, job := range jobs {
		time.Sleep(time.Second)
		fmt.Println(job, "by", w.color, "is done")
	}
}
