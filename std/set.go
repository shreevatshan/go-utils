package std

import "sync"

type Set struct {
	Elements map[interface{}]struct{}
	mutex    sync.Mutex
}

func InitSet() *Set {
	set := &Set{
		Elements: make(map[interface{}]struct{}),
	}
	return set
}

func (set *Set) Insert(element interface{}) {
	set.mutex.Lock()
	set.Elements[element] = struct{}{}
	set.mutex.Unlock()
}

func (set *Set) Delete(element interface{}) {
	set.mutex.Lock()
	delete(set.Elements, element)
	set.mutex.Unlock()
}

func (set *Set) Clear() {
	set.mutex.Lock()
	for k := range set.Elements {
		delete(set.Elements, k)
	}
	set.mutex.Unlock()
}

func (set *Set) Exists(element interface{}) bool {
	set.mutex.Lock()
	_, exists := set.Elements[element]
	set.mutex.Unlock()
	return exists
}

func (set *Set) Copy() *Set {
	set.mutex.Lock()
	newSet := InitSet()
	for element := range set.Elements {
		newSet.Elements[element] = struct{}{}
	}
	set.mutex.Unlock()
	return newSet
}
