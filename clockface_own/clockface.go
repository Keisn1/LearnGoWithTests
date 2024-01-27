package clockface

import (
	"io"
	"math"
	"text/template"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength float64 = 90
	minuteHandLength float64 = 80
	hourHandLength   float64 = 50
	clockCentreX     float64 = 150
	clockCentreY     float64 = 150
)

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := Round(math.Sin(angle), 0.05)
	y := Round(math.Cos(angle), 0.05)
	return Point{X: x, Y: y}
}

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY} //translate
	return p
}

func SecondsInRadians(t time.Time) float64 {
	return π / (30 / float64(t.Second()))
}

func DrawClockSecond(wr io.Writer, p Point) {
	tpl, err := template.ParseFiles("clock.xml")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(wr, p)
	if err != nil {
		panic(err)
	}
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
