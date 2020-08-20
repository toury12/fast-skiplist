package skiplist

import (
	"math/rand"
	"sync"
)

type Skey interface {
	Great(skey Skey) bool
	GreatE(skey Skey) bool
	Less(skey Skey) bool
	LessE(skey Skey) bool
}

type elementNode struct {
	next []*Element
}

type Element struct {
	elementNode
	key   Skey
	value interface{}
}

// Key allows retrieval of the key for a given Element
func (e *Element) Key() Skey {
	return e.key
}

// Value allows retrieval of the value for a given Element
func (e *Element) Value() interface{} {
	return e.value
}

// Next returns the following Element or nil if we're at the end of the list.
// Only operates on the bottom level of the skip list (a fully linked list).
func (e *Element) Next() *Element {
	return e.next[0]
}

type SkipList struct {
	elementNode
	maxLevel       int
	Length         int
	randSource     rand.Source
	probability    float64
	probTable      []float64
	mutex          sync.RWMutex
	prevNodesCache []*elementNode
}
