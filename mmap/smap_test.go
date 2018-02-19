package mmap

import "testing"

func assertPos(t *testing.T, pos *Position, x, y int64) {
	if pos.X != x || pos.Y != y {
		t.Fatal("invalid positoin", pos, x, y)
	}
}

func TestApi(t *testing.T) {
	smap := NewSMap(1)
	smap.AddTarget(1)
	smap.AddTarget(2)
	smap.AddTarget(3)
	smap.AddTarget(4)
	tar1 := smap.GetTarget(1)
	tar2 := smap.GetTarget(2)
	assertPos(t, tar1.Pos, 0, 0)

	smap.MoveOn(1, FRONT)
	// t1's position is 0,1
	assertPos(t, tar1.Pos, 0, 1)

	smap.MoveTo(2, newPosition(3, 8))
	// t2's position is 3,8
	assertPos(t, tar2.Pos, 3, 8)

	smap.MoveTo(3, newPosition(20, 21))

	smap.RemoveTarget(4)
	// mmap have target 1, 2
	mmap := smap.GetMMap(tar1.Pos)
	if len(mmap.Items) != 2 {
		t.Fatal("invalid mmap", len(mmap.Items))
	}
}
