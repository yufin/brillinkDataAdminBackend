package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TaskFactory struct {
}

func (TaskFactory) doTask(task int) (int, error) {
	// 模拟执行任务
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if rand.Intn(10) == 0 {
		return 0, fmt.Errorf("task %d failed", task)
	}
	return task + 1, nil
}

func (t TaskFactory) worker(wg *sync.WaitGroup, taskChan chan int, resultChan chan int, errChan chan error) {
	defer wg.Done()

	for task := range taskChan {
		result, err := t.doTask(task)
		if err != nil {
			errChan <- err
			return
		}
		fmt.Println("task done: ", task)

		resultChan <- result
	}
}

func (t TaskFactory) start() {
	taskChan := make(chan int, 100)
	resultChan := make(chan int, 100)
	errChan := make(chan error)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go t.worker(&wg, taskChan, resultChan, errChan)
	}

	// 发送任务
	go func() {
		for i := 0; i < 100; i++ {
			taskChan <- i
		}
		close(taskChan)
	}()

	done := make(chan bool)

	go func() {
		wg.Wait()
		close(done)
	}()

	// 等待结果或错误
	select {
	case <-done:
		fmt.Println("all workers done!")
		return
	//case result := <-resultChan:
	//	fmt.Printf("result: %d\n", result)
	case err := <-errChan:
		fmt.Printf("error: %v\n", err)
		return
	}
}
