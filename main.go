package main

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/abhishek-mane/go_parallel_stream_writer/concurrent/stream/writer"
)

func main() {

	w := writer.New(os.Stdout, 5)

	w.Add("Doing", "task1")
	w.Add("Doing", "task2")
	w.Add("Doing", "task3")
	w.Add("Doing", "work345")

	w.WriteInitial("Doing", "task1", "")
	w.WriteInitial("Doing", "task2", "")
	w.WriteInitial("Doing", "task3", "")
	w.WriteInitial("Doing", "work345", "")

	var wg sync.WaitGroup
	wg.Add(4)

	task := func(name, status string) {
		time.Sleep(time.Duration(10000000 * rand.Intn(100)))
		w.Write(name, status, "Done")
		wg.Done()
	}

	go task("Doing", "task1")
	go task("Doing", "task2")
	go task("Doing", "task3")
	go task("Doing", "work345")

	wg.Wait()

}
