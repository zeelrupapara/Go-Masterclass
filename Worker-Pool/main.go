package workerpool

import (
	"fmt"
	"sync"
	"time"
)

// Task represents a task to be processed
type Task struct {
	ID int
}

// WorkerPool processes tasks using a fixed number of workers
func WorkerPool(tasks []Task, numWorkers int) {
	taskChan := make(chan Task)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskChan, &wg)
	}

	// Send tasks to workers
	go func() {
		for _, task := range tasks {
			taskChan <- task
		}
		close(taskChan)
	}()

	// Wait for all workers to finish
	wg.Wait()
}

// Worker processes tasks from the task channel
func worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task.ID)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func main() {
	// Define tasks
	tasks := []Task{
		{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5},
		{ID: 6}, {ID: 7}, {ID: 8}, {ID: 9}, {ID: 10},
	}

	// Number of workers
	numWorkers := 3

	// Start the worker pool
	WorkerPool(tasks, numWorkers)
}
