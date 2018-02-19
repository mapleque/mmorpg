package mmap

import "math"

type Direction int

const (
	FRONT Direction = iota
	BACK
	LEFT
	RIGHT
)

type Position struct {
	X int64
	Y int64
}

// distance of two position
// Abs is absolute distance
// Trace is shortes trace distance
type Distance struct {
	Abs   int64
	Trace int64
}

type Scale struct {
	Origin *Position
	Radius int64
}

type Range struct {
	Start int64
	End   int64
}

func newPosition(x, y int64) *Position {
	return &Position{x, y}
}

func newScale(origin *Position, radius int64) *Scale {
	return &Scale{origin, radius}
}

func (this *Scale) positions() []*Position {
	ret := []*Position{}
	for x := this.Origin.X - this.Radius; x < this.Origin.X+this.Radius; x++ {
		for y := this.Origin.Y - this.Radius; y < this.Origin.Y+this.Radius; y++ {
			pos := newPosition(x, y)
			if pos.distanceTo(this.Origin) <= this.Radius {
				ret = append(ret, pos)
			}
		}
	}
	return ret
}

func (this *Position) distanceTo(pos *Position) int64 {
	return int64(math.Sqrt(math.Pow(float64(pos.X), 2) + math.Pow(float64(pos.Y), 2)))
}

func (this *Position) moveOn(dir Direction) *Position {
	switch dir {
	case FRONT:
		this.Y += 1
	case BACK:
		this.Y -= 1
	case LEFT:
		this.X += 1
	case RIGHT:
		this.X -= 1
	}
	return this
}

func (this *Position) moveTo(pos *Position) {
	this.X = pos.X
	this.Y = pos.Y
}
