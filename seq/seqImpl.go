package seq	

// Polymorph implementation of the ADT Sequence using a double-linked ringlist
// with administration block and dummy-element.

import "reflect"

type data struct {
	length  uint
	currIndex uint
	currElement *node
	anchor     *node
}

type node struct {
	content    Element
	next  *node
	previous *node
}

func New (e Element) *data {
	var sptr *data
	var nptr *node
	nptr = new (node)
	(*nptr).content = e // CAUTION: dirty for not native datatypes!! Actually deep copy necessary!! TODO
	(*nptr).next = nptr
	(*nptr).previous= nptr
	sptr = new (data)
	(*sptr).currElement = nptr
	(*sptr).anchor     = nptr
	return sptr
}

func (s *data)	Insert (e Element) {
	if reflect.TypeOf(e) == reflect.TypeOf ((*(*s).anchor).content) {
		var nptr *node
		nptr = new (node)										
		// set new node with its values ...
		(*nptr).content = e	// Dirty for not native datatypes - TODO
		(*nptr).next  = (*s).currElement                    
		(*nptr).previous = (*(*s).currElement).previous		
		// latch new node
		(*(*(*s).currElement).previous).next= nptr       
		(*(*s).currElement).previous = nptr			
		(*s).currIndex++											
		(*s).length++										
	}
}
	
func (s *data) Position (n uint) {
	if n < (*s).length {
		(*s).currIndex = n
	} else {
		(*s).currIndex = (*s).length
	}
	// set the pointer 'current'
	(*s).currElement = (*s).anchor
	for i:= uint(0);i <= (*s).currIndex;i++ {
		(*s).currElement = (*(*s).currElement).next
	}
}

func (s *data) CurrentIndex () uint {
	if (*s).currIndex == (*s).length {
		return (*s).length
	} else {
		return (*s).currIndex
	}
}

func (s *data) CurrentElement () (e Element, ok bool) {
	if (*s).currIndex < (*s).length {
		ok = true
		e = (*(*s).currElement).content // Dirty - TODO
	} else {
		ok = false
		e = (*(*s).anchor).content // Dirty for not native DT - TODO 
	}
	return
}

func (s *data) Forth () {
	if (*s).currIndex == (*s).length {
		return
	} else {
		(*s).currIndex++
		(*s).currElement = (*(*s).currElement).next
	}
}	

func (s *data) Back () {
	if (*s).currIndex == 0 {
		return
	} else {
		(*s).currIndex--
		(*s).currElement = (*(*s).currElement).previous
	}
}

func (s *data) Delete () {
	if (*s).length == (*s).currIndex { // no current element - nothing to do
		return
	} else { // There is a current element ...
		(*(*(*s).currElement).previous).next = (*(*s).currElement).next  
		(*(*(*s).currElement).next).previous = (*(*s).currElement).previous 
		(*s).currElement = (*(*s).currElement).next  
		(*s).length--  
	}
}

func (s *data) Length () uint {
	return (*s).length
}

