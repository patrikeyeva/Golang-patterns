package gofmt

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"unicode"
)

const (
	dot = "."
	tab = "\t"
)

type paragraph struct {
	idx   int
	value string
}

func run(ctx context.Context, fileName string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	input, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("error opening the file with text")
	}
	defer input.Close()

	output, err := os.OpenFile("output", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer output.Close()

	in := make(chan paragraph)
	go func() {
		defer close(in)
		read(ctx, input, in)
	}()

	out := make([]chan paragraph, runtime.NumCPU())
	for i := 0; i < len(out); i++ {
		out[i] = make(chan paragraph)
		go func(i int) {
			defer close(out[i])
			process(ctx, in, out[i])
		}(i)
	}

	write(ctx, output, merge(out))

	return ctx.Err()

}

func read(ctx context.Context, r io.Reader, out chan<- paragraph) {
	s := bufio.NewScanner(r)

	idx := 1
	for s.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			out <- paragraph{value: s.Text(), idx: idx}
			idx++
		}
	}
}

func process(ctx context.Context, in <-chan paragraph, out chan<- paragraph) {
	for text := range in {
		parts := strings.Fields(text.value)
		if len(parts) != 0 {
			parts[0] = tab + parts[0]
			for i := 1; i < len(parts); i++ {
				if len(parts[i]) > 0 && unicode.IsUpper([]rune(parts[i])[0]) && !strings.HasSuffix(parts[i-1], dot) {
					parts[i-1] += dot
				}
			}
			if !strings.HasSuffix(parts[len(parts)-1], dot) {
				parts[len(parts)-1] += dot
			}
		}

		out <- paragraph{idx: text.idx, value: strings.Join(parts, " ")}
	}
}

func merge(in []chan paragraph) chan paragraph {
	out := make(chan paragraph)

	var wg sync.WaitGroup

	wg.Add(len(in))

	for _, ch := range in {
		go func(ch chan paragraph) {
			defer wg.Done()
			for l := range ch {
				out <- l
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func write(ctx context.Context, w io.Writer, in <-chan paragraph) {
	nextIdx := 1
	m := make(map[int]paragraph)

	for p := range in {
		m[p.idx] = p

		for {
			if cached, ok := m[nextIdx]; ok {
				fmt.Fprintf(w, "%s\n", cached.value)
				delete(m, nextIdx)
				nextIdx++
			} else {
				break
			}
		}
	}

}
