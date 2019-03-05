package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"

	ct "github.com/daviddengcn/go-colortext"
)

type Outlet struct {
	Padding int

	sync.Mutex
}

var colors = []ct.Color{
	ct.Black,
	ct.Red,
	ct.Green,
	ct.Yellow,
	ct.Blue,
	ct.Magenta,
	ct.Cyan,
	ct.White,
}

func NewOutlet() *Outlet {
	return new(Outlet)
}

func (of *Outlet) LineReader(wg *sync.WaitGroup, name string, index int, r io.Reader, isError bool) {
	defer wg.Done()

	color := colors[index%len(colors)]
	reader := bufio.NewReader(r)

	var buffer bytes.Buffer
	for {
		buf := make([]byte, 1024)

		n, err := reader.Read(buf)
		if err != nil {
			return
		}
		buf = buf[:n]

		for {
			i := bytes.IndexByte(buf, '\n')
			if i < 0 {
				break
			}
			buffer.Write(buf[0:i])
			of.WriteLine(name, buffer.String(), color, ct.None, isError)
			buffer.Reset()
			buf = buf[i+1:]
		}

		buffer.Write(buf)
	}
}

func (of *Outlet) SystemOutput(str string) {
	of.WriteLine("dev-runner", str, ct.White, ct.None, false)
}

func (of *Outlet) ErrorOutput(str string) {
	fmt.Printf("ERROR: %s\n", str)
	os.Exit(1)
}

// Write out a single coloured line
func (of *Outlet) WriteLine(left, right string, leftC, rightC ct.Color, isError bool) {
	of.Lock()
	defer of.Unlock()

	ct.ChangeColor(leftC, true, ct.None, false)
	formatter := fmt.Sprintf("%%-%ds | ", of.Padding)
	fmt.Printf(formatter, left)

	ct.ResetColor()
	if isError {
		ct.ChangeColor(ct.Red, true, ct.None, true)
	}
	fmt.Println(right)
}
