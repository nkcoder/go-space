package topics

import (
	"fmt"
	"time"
)

// Think of the worker pool like a restaurant kitchen with multiple cooks.
// Instead of one cook doing all orders sequentially, you have several cooks working on different orders at the same time.

// each worker has an ID, receives jobs from a channel and send results to a channel
func poolWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		// process each job that comes through the channel
		fmt.Println("workder", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("workder", id, "finished job", j)
		// send result back
		results <- j * 2
	}
}

func WorkerPoolMain() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// start 3 workers and wait for jobs
	for w := 1; w <= 3; w++ {
		// We can pass a bidirectional channel to a function that expecting a unidirectional channel, Go automatically does the conversion.
		go poolWorker(w, jobs, results)
	}

	// distribute 5 jobs to the channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// tell works no more jobs are coming
	close(jobs)

	// read results from the channel
	for i := 1; i <= numJobs; i++ {
		<-results
	}

	// Each job taks 1 second, but with 3 workers, 5 jobs complete in ~2 seconds instead of 5.
	//$ time task run
	//workder 3 started job 1
	//workder 1 started job 2
	//workder 2 started job 3
	//workder 2 finished job 3
	//workder 2 started job 4
	//workder 1 finished job 2
	//workder 3 finished job 1
	//workder 1 started job 5
	//workder 1 finished job 5
	//workder 2 finished job 4
	//task run  0.25s user 0.37s system 21% cpu 2.823 total
}
