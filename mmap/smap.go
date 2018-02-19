package mmap

// Section map
// each entity has a sectionId
// many section map could build a world
//
// there are two index:
//  1. positionIndex
//  2. targetIndex
type SMap struct {
	SectionId     int64
	positionIndex _PositionIndex
	targetIndex   map[int64]*MapItem
}

// Api method from here //

// build a section map entity
func NewSMap(sectionId int64) *SMap {
	return &SMap{
		SectionId:     sectionId,
		positionIndex: newPositionIndex(),
		targetIndex:   make(map[int64]*MapItem),
	}
}

// add a target into section map
// with default position 0,0
func (this *SMap) AddTarget(target int64) {
	defaultPosition := newPosition(0, 0)
	this.addTarget(target, defaultPosition)
}

// remove a target from section map
func (this *SMap) RemoveTarget(target int64) {
	this.removeTarget(target)
}

// get a target map item pointer
// *MapItem @see mapitem.go
func (this *SMap) GetTarget(target int64) *MapItem {
	return this.getTarget(target)
}

// get a main map pointer
// which scale is a circle (pos,10)
// *Position @see position.go
// *MMap @see mmap.go
func (this *SMap) GetMMap(pos *Position) *MMap {
	scale := newScale(pos, 10)
	return this.getMMap(scale)
}

// target move step by step
// return the position pointer after move
// Direction @see position.go
// *Position @see position.go
func (this *SMap) MoveOn(target int64, dir Direction) *Position {
	return this.moveOn(target, dir)
}

// target move to a position directly
// *Position @see position.go
func (this *SMap) MoveTo(target int64, pos *Position) {
	this.moveTo(target, pos)
}

// Api method end //

//
func (this *SMap) addTarget(target int64, pos *Position) {
	item := newMapItem(target, pos)

	// add to position index
	this.positionIndex.add(item)
	// add to target index
	this.targetIndex[target] = item
}

func (this *SMap) removeTarget(target int64) {
	item, exist := this.targetIndex[target]
	if !exist {
		log.Warn("target %d can not be remove, does not exist", target)
		return
	}

	// remove from target index
	delete(this.targetIndex, target)

	// remove from position index
	this.positionIndex.remove(item)
}

func (this *SMap) getTarget(target int64) *MapItem {
	item, exist := this.targetIndex[target]
	if !exist {
		log.Error("target %d can not be found, does not exist", target)
		return nil
	}
	return item
}

func (this *SMap) getMMap(scale *Scale) *MMap {
	ret := &MMap{[]*MapItem{}}
	for _, pos := range scale.positions() {
		if itemMap, exist := this.positionIndex.get(pos); exist {
			for _, item := range itemMap {
				ret.Items = append(ret.Items, item)
			}
		}
	}
	return ret
}

func (this *SMap) moveOn(target int64, dir Direction) *Position {
	tar := this.getTarget(target)
	this.positionIndex.remove(tar)
	pos := tar.moveOn(dir)
	this.positionIndex.add(tar)
	return pos
}

func (this *SMap) moveTo(target int64, pos *Position) {
	tar := this.getTarget(target)
	this.positionIndex.remove(tar)
	tar.moveTo(pos)
	this.positionIndex.add(tar)
}

// [X][Y][target]
type _PositionIndex map[int64]map[int64]map[int64]*MapItem

func newPositionIndex() _PositionIndex {
	return make(map[int64]map[int64]map[int64]*MapItem)
}

func (this _PositionIndex) add(item *MapItem) {
	// build position index
	if _, exist := this[item.Pos.X]; !exist {
		this[item.Pos.X] = make(map[int64]map[int64]*MapItem)
	}
	if _, exist := this[item.Pos.X][item.Pos.Y]; !exist {
		this[item.Pos.X][item.Pos.Y] = make(map[int64]*MapItem)
	}

	this[item.Pos.X][item.Pos.Y][item.Target] = item
}

func (this _PositionIndex) remove(item *MapItem) {
	delete(this[item.Pos.X][item.Pos.Y], item.Target)

	// clear position index
	if len(this[item.Pos.X][item.Pos.Y]) == 0 {
		delete(this[item.Pos.X], item.Pos.Y)
	}
	if len(this[item.Pos.X]) == 0 {
		delete(this, item.Pos.X)
	}
}

func (this _PositionIndex) get(pos *Position) (map[int64]*MapItem, bool) {
	if x, existx := this[pos.X]; existx {
		if y, existy := x[pos.Y]; existy {
			return y, true
		}
	}
	return nil, false
}
