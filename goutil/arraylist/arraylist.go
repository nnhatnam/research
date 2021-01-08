package arraylist

import (
	"goutil"
)

func assertListImplementation() {
	var _ goutil.Queue = (*ArrayList)(nil)
}

type ArrayList struct {
	elements []interface{}
	len int
	modCount int
}



func NewArrayList() *ArrayList {
	return &ArrayList{
		elements: make([]interface{}, 10),
	}
}

func NewArrayListWithCap(cap int) *ArrayList {
	if cap < 0 {
		panic("Illegal Capacity")
	}
	return &ArrayList{
		elements: make([]interface{}, cap),
	}
}

func (list *ArrayList) Len() int {
	return list.len
}

func (list *ArrayList) Add(e interface{}) error {
	panic("implement me")
}

func (list *ArrayList) Element() (interface{}, error) {
	panic("implement me")
}

func (list *ArrayList) Offer() bool {
	panic("implement me")
}

func (list *ArrayList) Peek() interface{} {
	panic("implement me")
}

func (list *ArrayList) Poll() interface{} {
	panic("implement me")
}

func (list *ArrayList) Remove() (interface{}, error) {
	panic("implement me")
}

func (list *ArrayList) trimToSize() {
	list.modCount++
	list.resize(cap(list.elements))
}

func (list *ArrayList) ensureCapacityInternal(minCap int) {
	if minCap - len(list.elements) > 0 {
		list.grow(minCap)
	}
}

func (list *ArrayList) grow(minCap int) {
	oldCap := len(list.elements)
	newCap := oldCap + (oldCap >> 1)
	if newCap - oldCap < 0 {
		newCap = minCap
	}
	if newCap - _MAX_ARRAY_SIZE > 0 {
		newCap = list.hugeCapcity(minCap)
	}
	list.resize(newCap)
}

func (list *ArrayList) hugeCapcity(minCap int) int {
	if minCap < 0 { panic(errorString("Out of Memory"))}

	if minCap > _MAX_ARRAY_SIZE {
		return _MAX_ARRAY_SIZE
	}
	return minCap
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}