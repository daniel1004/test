package main

import "fmt"

type UnitType string

// Создайте 2 реализации интерфейса Dimensions: одна будет для дюймов, другая — для сантиметров.
//
//	Создайте по 1 реализации интерфейса Auto для автомобилей BMW, Mercedes и Dodge. Для реализации Dodge, dimensions должны возвращаться в дюймах.
const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	switch u.T {
	case Inch:
		if t == CM {
			return u.Value * 2.54
		}
		return u.Value
	case CM:
		if t == Inch {
			return u.Value / 2.54
		}
		return u.Value
	default:
		return u.Value
	}
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}
type dimentionsInCm struct {
	length, width, height Unit
}

func (d dimentionsInCm) Length() Unit {
	return d.length
}
func (d dimentionsInCm) Width() Unit {
	return d.width
}
func (d dimentionsInCm) Height() Unit {
	return d.height
}

type dimentionsInInch struct {
	length, width, height Unit
}

func (d dimentionsInInch) Length() Unit {
	return d.length
}
func (d dimentionsInInch) Width() Unit {
	return d.width
}
func (d dimentionsInInch) Height() Unit {
	return d.height
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}
type auto struct {
	brand, model         string
	dimensions           Dimensions
	maxSpeed, enginPower int
}

func (a auto) Brand() string {
	return a.brand
}
func (a auto) Model() string {
	return a.model
}
func (a auto) Dimensions() Dimensions {
	return a.dimensions
}
func (a auto) MaxSpeed() int {
	return a.maxSpeed
}
func (a auto) EnginePower() int {
	return a.enginPower
}

func main() {
	bmw := auto{
		"bmw",
		"x7",
		dimentionsInCm{
			Unit{100, CM},
			Unit{100, CM},
			Unit{100, CM},
		},
		290,
		180,
	}
	//fmt.Println(bmw.Dimensions().Length().Get(Inch))
	mercedes := auto{
		"Mercedes",
		"GLS",
		dimentionsInCm{
			Unit{111, CM},
			Unit{222, CM},
			Unit{333, CM},
		},
		280,
		360,
	}
	doge := auto{
		"Doge",
		"Shram",
		dimentionsInInch{
			Unit{230, CM},
			Unit{170, Inch},
			Unit{80, Inch},
		},
		180,
		450,
	}
	fmt.Println(bmw)
	fmt.Println(mercedes)
	fmt.Print(doge.model)
	fmt.Print("")
	fmt.Print(doge.dimensions)
	fmt.Print("")

}
