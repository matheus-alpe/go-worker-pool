package pool

import (
	"fmt"
	"math/rand"
	"time"
)

type Task interface {
	Process()
}

type EmailTask struct {
	Email string
}

func (e *EmailTask) Process() {
	fmt.Printf("@ Sending email to: %s\n", e.Email)
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Printf("@ -> Email sent to: %s\n", e.Email)
}

type ImageProcessTask struct {
	URL string
}

func (i *ImageProcessTask) Process() {
	fmt.Printf("# Processing the image: %s\n", i.URL)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	fmt.Printf("# -> Image processed: %s\n", i.URL)
}
