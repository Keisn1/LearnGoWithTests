package clockface_test

import (
	"GoWithTests/clockface"
	"bytes"
	"math"
	"testing"
	"time"
)

func TestSecondHandPoint(t *testing.T) {

	testCases := []struct {
		time time.Time
		want clockface.Point
	}{
		{
			time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			clockface.Point{X: 0, Y: 1},
		},
		{
			time.Date(1337, time.January, 1, 0, 0, 15, 0, time.UTC),
			clockface.Point{X: 1, Y: 0},
		},
		{
			time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC),
			clockface.Point{X: 0, Y: -1},
		},
		{
			time.Date(1337, time.January, 1, 0, 0, 45, 0, time.UTC),
			clockface.Point{X: -1, Y: 0},
		},
	}
	for _, tc := range testCases {
		got := clockface.SecondHandPoint(tc.time)
		if !roughlyEqualPoint(got, tc.want) {
			t.Errorf("clockface.SecondHand() = \"%v\"; want \"%v\"", got, tc.want)
		}
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := clockface.Point{X: 150, Y: 150 + 90}
	got := clockface.SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func roughlyEqual(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(p1, p2 clockface.Point) bool {
	return roughlyEqual(p1.X, p2.X) && roughlyEqual(p1.Y, p2.Y)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondsInRadians(c.time)
			if got != c.angle {
				t.Fatalf("Wanted %v radians, but got %v", c.angle, got)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func TestSecondHandPoint2(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockface.Point
	}{
		{simpleTime(0, 0, 30), clockface.Point{X: 0, Y: -1}},
		{simpleTime(0, 0, 45), clockface.Point{X: -1, Y: 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := clockface.SecondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Fatalf("Wanted %v Point, but got %v", c.point, got)
			}
		})
	}
}

func TestDrawClock(t *testing.T) {
	want := `<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">
<!-- bezel -->
  <circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>

<!-- hour hand -->
  <line x1="150" y1="150" x2="150" y2="100"
        style="fill:none;stroke:#ffd700;stroke-width:7px;"/>

 <!-- minute hand -->
  <line x1="150" y1="150" x2="150" y2="70"
        style="fill:none;stroke:#a9a9a9;stroke-width:7px;"/>

  <!-- second hand -->
  <line x1="150" y1="150" x2="60" y2="150"
        style="fill:none;stroke:#f00;stroke-width:3px;"/>
</svg>
`
	buffer := bytes.Buffer{}
	clockface.DrawClockSecond(&buffer, clockface.SecondHand(simpleTime(0, 0, 45)))
	if buffer.String() != want {
		t.Errorf("DrawClock() = \"%v\"; want \"%v\"", buffer.String(), want)
	}
}
