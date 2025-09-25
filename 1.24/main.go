package main

import "fmt"

// Расстояние между точками

// Разработать программу нахождения расстояния между двумя точками на плоскости.
// Точки представлены в виде структуры
// Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
// Расстояние рассчитывается по формуле между координатами двух точек.

// Подсказка: используйте функцию-конструктор NewPoint(x, y),
// Point и метод Distance(other Point) float64.

type IPoint interface {
	Distance(other IPoint) float64
}

type Point struct {
	x, y float64
}

func (p *Point) Distance(other IPoint) float64 {
	return (p.x + other.(*Point).x) * (p.y + other.(*Point).y)
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	p1 := NewPoint(1.0, 2.0)
	p2 := NewPoint(2.0, 4.0)

	distance := p1.Distance(p2)
	fmt.Println(distance)
}
