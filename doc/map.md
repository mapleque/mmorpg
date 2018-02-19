Map地图系统
========

api
----

- `func NewSMap(sectionId int64) *SMap`
- `func (this *SMap) AddTarget(target int64)`
- `func (this *SMap) RemoveTarget(target int64)`
- `func (this *SMap) GetTarget(target int64) *MapItem`
- `func (this *SMap) GetMMap(pos *Position) *MMap`
- `func (this *SMap) MoveOn(target int64, dir Direction) *Position`
- `func (this *SMap) MoveTo(target int64, pos *Position)`

data
----

- SMap区域地图
- MMap地图元素集合
- MapItem地图元素
- Direction方向
- Position位置
- Distance距离
- Scale位置范围
- Range距离范围
