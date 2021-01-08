package goutil

type Queue interface {

	//Add inserts the specified element into this queue, if it is possible to do so, immediately without violating
	//capacity restrictions, returning true upon success, and throwing an error if no space is currently available
	Add(e interface{}) (bool, error)

	//Element retrieves, but does not remove, the head of this queue. This method differs from "Peek" only in that it throws
	//an error if this queue is empty
	Element() (interface{}, error)

	//Offer inserts the specified element into this queue, if it is possible to do so, immediately without violating
	//capacity restrictions. When using capacity-restricted queue, this method is generally preferable to Add, which can
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


