package goutil

type LookupFunc func(e interface{}) bool

type Iterator interface {

	//HasNext returns true if the iteration has more elements.
	HasNext() bool

	//Next returns the next element in the iteration
	Next() interface{}

	//Remove removes from the underlying collection the last element returned by this iterator (optional operation).
	//This method can be called once per call. The behavior of an iterator is unspecified if the underlying collection
	//is modified while the iteration is in progress in any way other than by calling this method
	Remove() error
}

type Stack interface {

	//Empty tests if this stack is empty
	Empty() bool

	//Push pushes an item onto the top of this stack
	Push(e interface{})

	//Pop removes the item at the top of this stack and returns that item as the value of this function
	Pop() interface{}

	//Peek looks at the top of this stack without removing it from the stack
	Peek() interface{}
}

type Queue interface {

	//Add inserts the specified element into this queue, if it is possible to do so, immediately without violating
	//capacity restrictions, returning nil upon success, and throwing an error if no space is currently available
	Add(e interface{}) error

	//Element retrieves, but does not remove, the head of this queue. This method differs from "Peek" only in that it throws
	//an error if this queue is empty
	Element() (interface{}, error)

	//Offer inserts the specified element into this queue, if it is possible to do so, immediately without violating
	//capacity restrictions. Returning true upon success and false if no space is currently available.
	//When using a capacity-restricted queue, this method is generally preferable to Add, which can
	//fail to insert an element only by throwing an exception
	Offer() bool

	//Peek retrieves, but does not remove, the head of this queue, or return nil if the queue is empty
	Peek() interface{}

	//Pool retrieves and remove the head of this queue, or return nil if the queue is empty
	Poll() interface{}

	//Remove retrieves and remove the head of this queue. This method differs from Poll only in that it throws an error
	//if this queue is empty.
	Remove() (interface{}, error)
}


type Deque interface {
	Queue
	Stack

	//AddFirst inserts the specified element at the front of this deque, if it is possible to do so, immediately without violating
	//capacity restrictions. When using a capacity-restricted deque, it is generally preferable to use method OfferFirst
	AddFirst(e interface{}) error

	//AddLast inserts the specified element at the of this this deque, if it is possible to do so, immediately without
	//violating capacity restrictions. When using a capacity-restricted deque, it is generally preferable to use method OfferLast
	AddLast(e interface{}) error

	//Contains applies function f to each element in this deque, return true if f return true in any element.
	Contains(f LookupFunc) bool

	//DescendingIterator returns an iterator over the elements in this deque in reverse sequential order. The elements will
	//be returned in order from last (tail) to first (head)
	DescendingIterator() Iterator

	//Iterator returns an iterator over the elements in this deque in proper sequence. The elements will be returned in
	//order from first (head) to last(tail)
	Iterator() Iterator

	//First retrieves, but not removes, the first element of this deque. Return an error if the deque is empty
	First() (interface{}, error)

	//Last retrieves, but not removes, the last element of this deque. Return an error if the deque is empty
	Last() (interface{}, error)

	//OfferFirst inserts the specified element at the front of the queue represented by this deque (in other words, at the tail of
	//this deque) if it is possible to do so immediately without violating capacity restrictions. When using a capacity-restricted
	//deque, this method is generally preferable to the AddFirst method, which can fail to insert an element only by throwing an exception
	OfferFirst(e interface{}) bool

	//OfferLast inserts the specified element at the end of this deque unless it would violate capacity restrictions. When using
	//a capacity-restricted deque, this method is generally preferable to the AddLast method, which can fail to insert an element
	//only by throwing an exception
	OfferLast(e interface{}) bool

	//PeekFirst retrieves, but not removes, the first element of this deque. Return nil if the deque is empty
	PeekFirst() interface{}

	//PeekLast retrieves, but not removes, the last element of this deque. Return nil if the deque is empty
	PeekLast() interface{}

	//PollFirst retrieves and removes the first element of this deque or return nil if the deque is empty
	PollFirst() interface{}

	//PollFirst retrieves and removes the last element of this deque or return nil if the deque is empty
	PollLast() interface{}

	//Push pushes an element onto the stack represented by this deque (in other words, at the head of this deque) if it
	//is possible to do so immediately without violating capacity restrictions, returning nil upon success and throwing
	//an error if no space is currently available.
	//Push(e interface{}) error

	//RemoveFirst removes the first element of this deque. Throw an error if the deque is empty
	RemoveFirst() error

	//RemoveFirst removes the last element of this deque. Throw an error if the deque is empty
	RemoveLast() error

	//Return the numbers of elements in this deque
	Len() int

}
