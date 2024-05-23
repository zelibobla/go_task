package orderedmap

import (
	"container/list"
	"sync"
)

type entry struct {
	key     string
	value   string
	element *list.Element
}

type OrderedMap struct {
	mu    sync.RWMutex
	items map[string]*entry
	order *list.List
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		items: make(map[string]*entry),
		order: list.New(),
	}
}

func (om *OrderedMap) Add(key, value string) {
	om.mu.Lock()
	defer om.mu.Unlock()
	if e, exists := om.items[key]; exists {
		e.value = value
	} else {
		elem := om.order.PushBack(key)
		om.items[key] = &entry{key, value, elem}
	}
}

func (om *OrderedMap) Delete(key string) {
	om.mu.Lock()
	defer om.mu.Unlock()
	if e, exists := om.items[key]; exists {
		om.order.Remove(e.element)
		delete(om.items, key)
	}
}

func (om *OrderedMap) Get(key string) (string, bool) {
	om.mu.RLock()
	defer om.mu.RUnlock()
	if e, exists := om.items[key]; exists {
		return e.value, true
	}
	return "", false
}

func (om *OrderedMap) GetAll() map[string]string {
	om.mu.RLock()
	defer om.mu.RUnlock()
	result := make(map[string]string, len(om.items))
	for e := om.order.Front(); e != nil; e = e.Next() {
		key := e.Value.(string)
		result[key] = om.items[key].value
	}
	return result
}
