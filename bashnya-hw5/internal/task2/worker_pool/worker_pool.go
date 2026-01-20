package worker_pool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	workers int
	jobs    <-chan any
	ctx     context.Context
	wg      *sync.WaitGroup
}

func NewWorkerPool(workers int, jobs <-chan any, ctx context.Context) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		jobs:    jobs,
		ctx:     ctx,
		wg:      &sync.WaitGroup{},
	}
}

func (p *WorkerPool) Start() {
	p.wg.Add(p.workers)

	for i := 0; i < p.workers; i++ {
		go func(id int) {
			defer p.wg.Done()
			p.worker(id)
		}(i)
	}
}

func (p *WorkerPool) worker(id int) {
	for {
		select {
		case <-p.ctx.Done():
			fmt.Printf("Worker %d canceled\nExit...\n", id)
			return
		case job, ok := <-p.jobs:
			if !ok {
				fmt.Printf("Worker %d canceled\nExit...\n", id)
			} else {
				fmt.Printf("worker %d: %v process...\n", id, job)
			}
			return
		}
	}
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func GenerateWorkers(ctx context.Context, jobs chan<- any) {
	i := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Generating canceled\nExit...\n")
			return
		default:
			jobs <- fmt.Sprintf("job #%d", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}
