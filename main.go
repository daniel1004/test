package main

import (
	"container/ring"
	"fmt"
)

// Вы должны написать код, который динамически рассчитывает
// среднюю температуру в городе от 10 метеостанций.
// Данные от каждой станции постоянно обновляются.
// Представьте, что событие изменения температуры
// от какой-либо метеостанции — это просто ввод соответствующей температуры
// через консоль и номера метеостанции,
// от которой она получена.
// При каждом изменении температуры на любой станции
// ваша программа должна сразу вывести среднюю температуру в городе.
// (кольцевой массив)
//1) СОЗДАТЬ МЕТЕОСТАНЦИИ
//2) ИНИЦИАЛИЗИРОВАТЬ ИХ
//3) ПОСЧИТАТЬ СРЕДНЮЮ
//4) БЕСКОНЕЧНЫЙ ЦИКЛ НА ИЗМЕНЕНИЕ ТЕМПЕРАТУРЫ
//5) ВЫЗОВ ИЗМЕНЕННЫХ ПОКАЗАТЕЛЕЙ

type MeteoStation struct {
	Name int
	temp float64
}

func main() {
	// СОЗДАТЬ МЕТЕОСТАНЦИИ
	NewRing := ring.New(10)

	// ИНИЦИАЛИЗИРОВАТЬ ИХ
	for i := 1; i < NewRing.Len()+1; i++ {
		NewRing.Value = MeteoStation{
			i,
			0.0,
		}
		NewRing = NewRing.Next()
	}
	//ПОСЧИТАТЬ СРЕДНЮЮ
	sum := 0.0
	calcSrZnach := func(*ring.Ring) float64 {
		NewRing.Do(func(p any) {
			sum += p.(MeteoStation).temp
		})
		return sum / float64(NewRing.Len())
	}
	for {
		fmt.Println("Введите номер метеостанции ")
		var statioNnumber int
		fmt.Scan(&statioNnumber)
		for statioNnumber < 0 || statioNnumber > NewRing.Len() {
			fmt.Println("Такой станции нет, введите номер от 0 до 10 ")
			fmt.Scan(&statioNnumber)
		}
		fmt.Println("Введите температуру ")
		var newTemp float64
		fmt.Scan(&newTemp)

		// НЕ ПОНИИМАЮ ПОЧЕМУ НЕ РАБОТАЕТ
		//NewRing.Do(func(p any) {
		//	if p != nil {
		//		ms, ok := p.(MeteoStation)
		//		if ok {
		//			if statioNnumber == ms.Name {
		//				ms = MeteoStation{
		//					Name: statioNnumber,
		//					temp: newTemp,
		//				}
		//			}
		//		}
		//	}
		//})
		current := NewRing
		for i := 0; i < statioNnumber; i++ {
			current = current.Next()
		}
		current.Value = MeteoStation{
			Name: statioNnumber,
			temp: newTemp,
		}
		result := calcSrZnach(NewRing)
		fmt.Println(result)
	}
}
