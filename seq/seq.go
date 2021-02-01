package seq

/* A sequence is a list of ongoing nummerated objects (of the same
 * datatype). The same object can be contained mulitple times. The
 * index of the first object is 0!
 * 
 * As objects of the sequence we consider arbitrary instances of a 
 * datatype as the case may be of an abstract datatype. */
type Element interface {}

// Requirements: -
// Result: a new, empty sequence of objects of the 'real' datatype of
//         e, called: 'Elementype of the sequence'.
// New (e Element) Folge

type Sequence interface {
// Requirements: -
// Effect: If e wasn't of the elementtype of the sequence, nothing has
//		   happened. Otherwise: e is inserted into the sequence in front of
//		   the current element. If there wasn't a current element, e
//		   is inserted at the end of the sequence
	Insert(e Element)
	
// Req.: -
// Eff.: If there was an element with index n in the sequence, it is 
//		 now the current element. Otherwise there is no current 
//		 element.
	Position (n uint)
	
// Req.: -
// Result: If the sequence was empty or there was no current element
//		   the length of the sequence is yielded, otherwise the 
//		   index of the position of the current element is returned.
	CurrentIndex() uint
	
// Req.: -
// Res.: If there was no current element, ok is false and e is an 
//		 instance with the type of the elements of the sequence, 
//		 which itself isn't part of the sequence of objects. Otherwise
//		 ok is true and the current element e is returned.
	CurrentElement() (e Element, ok bool)
	
// Req.: -
// Eff.: If there was no current element nothing has happened. 
//		 Otherwise the element which follows the previously current
//		 element now is current. If there is no such element, no 
//		 element is current.
	Forth() 

// Req.: -
// Eff.: If the element with index 0 was current, nothing has happened
//		 If there was no current element, the last element is current.
//		 Otherwise the element which precedes the previously current
//		 element now is current.
	Back()

// Req.: -
// Eff.: If there was a current element, it is deleted from the
//		 sequence. The rest stays in the same order. The previously
//		 following element now is current. If there is no such element
//		 no element is current. If there was no current element, nothing
//		 has happened.
	Delete()
	
// Req.: -
// Res.: The length of the sequence, i.e. the number of elements in 
//		 the sequence, is returned.
	Length() uint
}

