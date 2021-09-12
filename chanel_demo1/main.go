package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	value int64
}
type JobResult struct {
	job *Job
	sum int64
}

func producter(p chan<- *Job) {
	wg.Done()
	for {
		x := rand.Int63()
		newJob := &Job{
			value: x,
		}
		p <- newJob
		time.Sleep(time.Second)
	}
}

func consumer(p <-chan *Job, result chan<- *JobResult) {
	wg.Done()
	for {
		job := <-p
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &JobResult{
			job: job,
			sum: sum,
		}
		result <- newResult
	}
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *JobResult, 100)
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go producter(jobChan)

	wg.Add(24)
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d sum: %d\n", result.job.value, result.sum)
	}
	wg.Wait()

}
