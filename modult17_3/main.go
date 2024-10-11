package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var bufferSize int = 3

const bufferDrainInterval time.Duration = 20 * time.Second

type Buffer struct {
	position int
	arr      []int
	size     int
	mu       sync.Mutex
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		position: -1,
		arr:      make([]int, size),
		size:     size,
	}
}

func (b *Buffer) Push(el int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.position == b.size-1 {
		for i := 1; i <= b.size-1; i++ {
			b.arr[i-1] = b.arr[i]
		}
		b.arr[b.position] = el
	} else {
		b.position++
		b.arr[b.position] = el
	}

}
func (b *Buffer) Pop() []int {
	if b.position == -1 {
		return nil
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	output := b.arr[:b.position+1]
	b.position = -1
	return output
}
func FilterPositive(done <-chan struct{}, inputData <-chan int) <-chan int {
	onlyPositiveData := make(chan int)
	go func() {
		defer close(onlyPositiveData)
		for {
			select {
			case <-done:
				return
			case value, ok := <-inputData:
				if !ok {
					return
				}
				if value > 0 {
					select {
					case onlyPositiveData <- value:
					case <-done:
						return
					}
				}
			}
		}
	}()
	return onlyPositiveData
}

func FilterThree(done <-chan struct{}, inputData <-chan int) <-chan int {
	onlyThreeData := make(chan int)
	go func() {
		defer close(onlyThreeData)
		for {
			select {
			case <-done:
				return
			case value, ok := <-inputData:
				if !ok {
					return
				}
				if value%3 == 0 {
					select {
					case onlyThreeData <- value:
					case <-done:
					}
				}

			}
		}
	}()
	return onlyThreeData
}

func main() {
	inputData := func() (<-chan struct{}, <-chan int) {
		output := make(chan int)
		done := make(chan struct{})
		go func() {
			defer close(done)
			scanner := bufio.NewScanner(os.Stdin)
			var str string
			fmt.Println("Press enter to continue")
			for {
				scanner.Scan()
				str = scanner.Text()
				if strings.EqualFold(str, "stop") {
					fmt.Println("Программа завершила работу")
					return
				}
				val, ok := strconv.Atoi(scanner.Text())
				if ok != nil {
					fmt.Println("Только int!")
					continue
				}
				output <- val
			}
		}()
		return done, output
	}

	buferisation := func(done <-chan struct{}, input <-chan int) <-chan int {
		buffer := NewBuffer(bufferSize)
		bufferedIntChan := make(chan int)

		go func() {
			defer close(bufferedIntChan)
			for {
				select {
				case <-done:
					return
				case value, ok := <-input:
					if !ok {
						return
					}
					buffer.Push(value)
				}
			}
		}()

		go func() {
			for {
				select {
				case <-done:
					close(bufferedIntChan)
					return
				case <-time.After(bufferDrainInterval):
					bufferData := buffer.Pop()
					for _, val := range bufferData {
						select {
						case bufferedIntChan <- val:
						case <-done:
						}
					}
				}
			}
		}()
		return bufferedIntChan
	}
	potrebitel := func(done <-chan struct{}, input <-chan int) {
		for {
			select {
			case <-done:
				return
			case val := <-input:
				fmt.Printf("Обработаны данные: %d\n", val)
			}
		}
	}

	done, output := inputData()
	potrebitel(done, buferisation(done, FilterThree(done, FilterPositive(done, output))))
}
