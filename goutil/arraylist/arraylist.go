package arraylist

import (
	"goutil"
)

//func assertListImplementation() {
//	var _ goutil.Queue = (*ArrayList)(nil)
//}

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

func (list *ArrayList) IsEmpty() bool {
	return list.len == 0
}

func (list *ArrayList) Add(e interface{}) error {
	//TO DO: Is it neccessary to do range check ?
	list.ensureCapacityInternal(list.len + 1) //increments modCount++
	list.elements[list.len] = e
	list.len++
	return nil
}

func (list *ArrayList) AddAt(index int, e interface{}) error {
	list.rangeCheckForAdd(index)

	list.ensureCapacityInternal(list.len + 1) //increments modCount++\
	copy(list.elements[index+1:], list.elements[index:])
	list.elements[index] = e
	list.len++
	return nil
}

func (list *ArrayList) Element() (interface{}, error) {
	if list.len > 0 {
		return list.elements[0], nil
	}
	return nil, errorString("List is empty")
}

func (list *ArrayList) Offer(e interface{}) bool {
	panic("implement me")
}

func (list *ArrayList) Peek() interface{} {
	panic("implement me")
}

func (list *ArrayList) Poll() interface{} {
	panic("implement me")
}

func (list *ArrayList) RemoveFunc(f goutil.LookupFunc) interface{} {
	panic("implement me")
}

func (list *ArrayList) AddAll(f goutil.LookupFunc) interface{} {
	panic("implement me")
}


func (list *ArrayList) Remove(index int) interface{} {
	list.rangeCheck(index)

	list.modCount++
	oldValue := list.elements[index]
	numMoved := list.len - index - 1
	if numMoved > 0 {
		copy(list.elements[index:] , list.elements[index+1:])
		list.elements[index] = nil
		list.len--
	}
	return oldValue
}

func (list *ArrayList) fastRemove(index int) {
	list.modCount++
	numMoved := list.len - index - 1
	if numMoved > 0 {
		copy(list.elements[index:] , list.elements[index+1:])
		list.elements[index] = nil
		list.len--
	}
}

func (list *ArrayList) Get(index int) interface{} {
	list.rangeCheck(index)
	return list.elements[index]
}

func (list *ArrayList) Set(index int, element interface{})  {
	list.rangeCheck(index)
	list.elements[index] = element
}


func (list *ArrayList) Contains(f goutil.LookupFunc) bool {
	return list.IndexOf(f) >= 0
}

func (list *ArrayList) Clone() *ArrayList {
	newElements := make([]interface{}, len(list.elements))
	copy(newElements, list.elements)
	return &ArrayList{
		elements: newElements,
	}
}

func (list *ArrayList) ToArray() []interface{} {
	newElements := make([]interface{}, len(list.elements))
	copy(newElements, list.elements)
	return newElements
}



func (list *ArrayList) IndexOf(f goutil.LookupFunc) int {
	for k, v := range list.elements {
		if f(v) {
			return k
		}
	}
	return -1
}

func (list *ArrayList) LastIndexOf(f goutil.LookupFunc) int {
	for i := len(list.elements) - 1; i >= 0; i-- {
		if f(list.elements[i]) {
			return i
		}
	}

	return -1
}

func (list *ArrayList) rangeCheck(index int) {
	if index < 0 || index >= list.len {
		panic(errorString("Index out of range"))
	}
}

func (list *ArrayList) rangeCheckForAdd(index int) {
	if index < 0 || index > list.len {
		panic(errorString("Index out of range"))
	}
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
		newCap = list.hugeCapacity(minCap)
	}
	list.resize(newCap)
}

func (list *ArrayList) hugeCapacity(minCap int) int {
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