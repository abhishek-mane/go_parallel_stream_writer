package writer

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// ParallelStreamWriter :
type ParallelStreamWriter struct {
	lines  []string
	width  int
	stream *os.File
	mux    sync.Mutex
}

// Add :
func (w *ParallelStreamWriter) Add(operation, entity string) {

	id := fmt.Sprintf("%s %s", operation, entity)
	w.lines = append(w.lines, id)
	w.width = func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}(w.width, len(id))

}

// WriteInitial :
func (w *ParallelStreamWriter) WriteInitial(operation, entity, status string) {

	str := "%-" + strconv.Itoa(w.width) + "v ... %s\r\n"
	w.stream.Write([]byte(fmt.Sprintf(str, fmt.Sprintf("%s %s", operation, entity), status)))
}

// Write :
func (w *ParallelStreamWriter) Write(operation, entity, status string) {

	id := fmt.Sprintf("%s %s", operation, entity)

	w.mux.Lock()
	position := getIndex(w.lines, operation+" "+entity)
	diff := len(w.lines) - position

	// move up
	w.stream.Write([]byte(fmt.Sprintf("%c[%dA", 27, diff)))

	// erase
	w.stream.Write([]byte(fmt.Sprintf("%c[2K\r", 27)))
	str := "%-" + strconv.Itoa(w.width) + "v ... %s\r"
	w.stream.Write([]byte(fmt.Sprintf(str, id, status)))

	// move back down
	w.stream.Write([]byte(fmt.Sprintf("%c[%dB", 27, diff)))
	w.mux.Unlock()

}

// New :
func New(stream *os.File, len int) *ParallelStreamWriter {

	w := ParallelStreamWriter{stream: stream}
	w.lines = make([]string, len)
	return &w

}

func getIndex(slice []string, val string) int {
	for idx, element := range slice {
		if element == val {
			return idx
		}
	}
	return 0
}
