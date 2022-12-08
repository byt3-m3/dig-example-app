package worker

import "context"

type worker struct {
}

func New() *worker {
	return &worker{}
}

func (w worker) Run(ctx context.Context) error {

	return nil
}
