package maps

import (
	"sync"
)

type Set struct {
	elements map[interface{}]struct{}
	sync.RWMutex
}

func NewSet() *Set {
	set := &Set{
		elements: make(map[interface{}]struct{}),
	}
	return set
}

func ToSet(elements []interface{}) *Set {
	set := NewSet()
	for _, element := range elements {
		// directly insert the element into the set
		// lock is not required as the set is not shared yet
		set.SafeInsert(element)
	}
	return set
}

func (set *Set) Insert(element interface{}) {
	set.Lock()
	defer set.Unlock()
	set.elements[element] = struct{}{}
}

// SafeInsert is a thread-safe version of Insert
// It does not lock the set as it assumes that the caller has already locked it
func (set *Set) SafeInsert(element interface{}) {
	set.elements[element] = struct{}{}
}

func (set *Set) Delete(element interface{}) {
	set.Lock()
	defer set.Unlock()
	delete(set.elements, element)
}

func (set *Set) Clear() {
	set.Lock()
	defer set.Unlock()
	for k := range set.elements {
		delete(set.elements, k)
	}
}

func (set *Set) Exists(element interface{}) bool {
	set.RLock()
	defer set.RUnlock()
	_, exists := set.elements[element]
	return exists
}

// SafeExists is a thread-safe version of Exists
// It does not lock the set as it assumes that the caller has already locked it
func (set *Set) SafeExists(element interface{}) bool {
	_, exists := set.elements[element]
	return exists
}

func (set *Set) Size() int {
	set.RLock()
	defer set.RUnlock()
	return len(set.elements)
}

func (set *Set) List() []interface{} {
	set.RLock()
	defer set.RUnlock()
	list := make([]interface{}, 0, len(set.elements))
	for k := range set.elements {
		list = append(list, k)
	}
	return list
}

func (set *Set) Copy() *Set {
	set.Lock()
	defer set.Unlock()
	newSet := NewSet()
	for element := range set.elements {
		newSet.elements[element] = struct{}{}
	}
	return newSet
}
