package job

import (
	"context"
	"sync"

	"github.com/user/finder-clone/internal/ops"
)

// Queue manages background file operations
type Queue struct {
	jobs    chan Job
	results chan ops.Progress
	wg      sync.WaitGroup
	workers int
}

type Job interface {
	Execute(ctx context.Context, progress chan<- ops.Progress) error
}

func NewQueue(workers int) *Queue {
	return &Queue{
		jobs:    make(chan Job, 100),
		results: make(chan ops.Progress, 100),
		workers: workers,
	}
}

func (q *Queue) Start(ctx context.Context) {
	for i := 0; i < q.workers; i++ {
		q.wg.Add(1)
		go q.worker(ctx)
	}
}

func (q *Queue) worker(ctx context.Context) {
	defer q.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-q.jobs:
			// Execute job and pipe progress
			// In a real app, we'd handle job IDs and per-job progress channels
			job.Execute(ctx, q.results)
		}
	}
}

func (q *Queue) Submit(j Job) {
	q.jobs <- j
}
