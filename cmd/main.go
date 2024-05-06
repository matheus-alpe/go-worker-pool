package main

import (
	"fmt"

	"github.com/matheus-alpe/go-worker-pool/internal/pool"
)

func main() {
	tasks := []pool.Task{
		&pool.EmailTask{Email: "matttalves@gmail.com"},
		&pool.ImageProcessTask{URL: "image.png"},
		&pool.EmailTask{Email: "test@gmail.com"},
		&pool.ImageProcessTask{URL: "random-image.png"},
		&pool.EmailTask{Email: "a@gmail.com"},
		&pool.ImageProcessTask{URL: "tree-image.png"},
		&pool.EmailTask{Email: "b@yahoo.com"},
		&pool.ImageProcessTask{URL: "leaf-image.png"},
	}

	wp := pool.NewWorker(tasks, 5)
	wp.Run()
	fmt.Println("All tasks have been processed")
}
