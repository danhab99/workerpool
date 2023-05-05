package main

import "sync"

type Workerpool struct {
	tasks chan func()
	wg    *sync.WaitGroup
}

func New(count int) (out Workerpool) {
	out.tasks = make(chan func(), 1e10)

	for i := 0; i < count; i++ {
		go func() {
			defer out.wg.Done()
			for f := range out.tasks {
				f()
			}
		}()
	}

	return
}

func (pool Workerpool) Add(f func()) {
	pool.tasks <- f
}

func (pool Workerpool) Stop() {
	close(pool.tasks)
	pool.wg.Done()
}
