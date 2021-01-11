package arraylist

import "goutil"

func assertIteratorImplementation() {
	var _ goutil.Iterator = (*hiddenItr)(nil)
}


type hiddenItr struct {
	cursor int
	lastRet int
	expectedModCount int
	list *ArrayList
}

func (h *hiddenItr) HasNext() bool {
	return h.cursor ==  h.list.len
}

func (h *hiddenItr) Next() interface{} {
	h.checkForComodification()
	i := h.cursor
	if i >= h.list.len {
		panic(errorString("NoSuchElementException"))
	}
	if i >= h.list.len {
		panic(errorString("ConcurrentModificationException"))
	}
	h.cursor = i+1
	h.lastRet = i
	return h.list.elements[i]
}

func (h *hiddenItr) Remove() error {
	if h.lastRet < 0 {
		panic(errorString("IllegalStateException"))
	}
	h.checkForComodification()
	h.list.Remove(h.lastRet)
	h.cursor = h.lastRet
	h.lastRet = -1
	h.expectedModCount = h.list.modCount
	return nil
}

func (h *hiddenItr) checkForComodification() {
	if h.list.modCount != h.expectedModCount {
		panic(errorString("ConcurrentModificationException"))
	}
}
