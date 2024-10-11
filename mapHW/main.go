package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var bufferSize int = 3

const bufferDrainInterval time.Duration = 10 * time.Second

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
					close(output)
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
		r := ring.New(bufferSize)
		preR := r
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
					r.Value = value
					r = r.Next()

				}
			}
		}()

		go func() {
			for {
				select {
				case <-done:
					return
				case <-time.After(bufferDrainInterval):

					preR.Do(func(p interface{}) {
						if p != nil {
							select {
							case bufferedIntChan <- p.(int):
							case <-done:
								return
							}

							preR.Value = nil
							preR = preR.Next()
						}
					})
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
