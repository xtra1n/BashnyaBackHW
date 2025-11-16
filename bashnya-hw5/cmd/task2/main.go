package main

import (
	"bashnya-hw5/internal/task2/cli"
	"bashnya-hw5/internal/task2/worker_pool"
	io_utils "bashnya-hw5/pkg/utils"
)

func main() {
	workers, err := cli.ParseFlags()

	if err != nil {
		io_utils.PrintError(err)
		return
	}

	ctx, cancel := cli.ProcessSignal()
	defer cancel()

	jobs := make(chan any)

	go worker_pool.GenerateWorkers(ctx, jobs)

	pool := worker_pool.NewWorkerPool(workers, jobs, ctx)
	pool.Start()

	pool.Wait()

}
