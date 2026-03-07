package main

import (
	"fmt"
	"time"
)

// Task — структура задания
type Task struct {
	id   int
	name string
}

// WorkerPool управляет рабочими и каналом заданий
type WorkerPool struct {
	tasks chan Task
	done  chan bool
	size  int
}

// NewWorkerPool создаёт новый пул рабочих с заданным количеством рабочих
func NewWorkerPool(size int) *WorkerPool {
	wp := &WorkerPool{
		tasks: make(chan Task),
		done:  make(chan bool),
		size:  size,
	}
	return wp
}

// Start запускает пулы рабочих
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.size; i++ {
		go func(workerId int) {
			for task := range wp.tasks {
				fmt.Printf("Worker #%d processing task '%s'\n", workerId, task.name)
				time.Sleep(time.Duration(task.id) * time.Second) // Имитация нагрузки
				fmt.Printf("Task '%s' completed by worker #%d\n", task.name, workerId)
			}
			wp.done <- true
		}(i)
	}
}

// Submit отправляет задание в канал
func (wp *WorkerPool) Submit(task Task) {
	wp.tasks <- task
}

// Stop останавливает пул рабочих
func (wp *WorkerPool) Stop() {
	close(wp.tasks)
	for i := 0; i < wp.size; i++ {
		<-wp.done
	}
}

func main() {
	poolSize := 3   // Количество рабочих
	taskCount := 10 // Количество заданий

	wp := NewWorkerPool(poolSize)
	wp.Start()

	for i := 0; i < taskCount; i++ {
		task := Task{id: i, name: fmt.Sprintf("task-%d", i)}
		wp.Submit(task)
	}

	wp.Stop()
	fmt.Println("Все задания выполнены")
}
