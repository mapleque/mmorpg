package mmap

type MapItem struct {
	Target int64
	Pos    *Position
}

func newMapItem(target int64, pos *Position) *MapItem {
	return &MapItem{
		Target: target,
		Pos:    pos,
	}
}

func (this *MapItem) moveOn(dir Direction) *Position {
	return this.Pos.moveOn(dir)
}

func (this *MapItem) moveTo(pos *Position) {
	this.Pos.moveTo(pos)
}
