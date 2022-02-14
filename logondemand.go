package logondemand

import (
	"io"
	"os"
	"syscall"
	"time"
)

// New create fifo if nessesary, which can be used in `log.New()`
func New(filename string) (io.WriteCloser, error) {
	// mkfifo
	err := syscall.Mkfifo("/tmp/"+filename, 0755)
	if !os.IsExist(err) && err != nil {
		return nil, err
	}
	// open
	l, err := os.OpenFile("/tmp/"+filename, os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Cat cat fifo content like shell `cat`. Restart after io.EOF.
func Cat(filename string) error {
	l, err := os.Open("/tmp/" + filename)
	if err != nil {
		return err
	}
	for err != io.EOF {
		_, err = io.Copy(os.Stderr, l)
		time.Sleep(500 * time.Millisecond)
	}
	return err
}
