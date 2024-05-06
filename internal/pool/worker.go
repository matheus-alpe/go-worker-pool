package pool

import "sync"

type Worker struct {
	Tasks            []Task
	concurrencyLevel int
	tasksChan        chan Task
	wg               sync.WaitGroup
}

func NewWorker(tasks []Task, concurrencyLevel int) *Worker {
	return &Worker{
		Tasks:            tasks,
		concurrencyLevel: concurrencyLevel,
	}
}

func (w *Worker) work() {
	for task := range w.tasksChan {
		task.Process()
		w.wg.Done()
	}
}

func (w *Worker) Run() {
	l := len(w.Tasks)
	w.tasksChan = make(chan Task, l)
	w.wg.Add(l)

	for i := 0; i < w.concurrencyLevel; i++ {
		go w.work()
	}

	for _, task := range w.Tasks {
		w.tasksChan <- task
	}
	close(w.tasksChan)

	w.wg.Wait()
}
