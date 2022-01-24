package main

import (
	"fmt"
	"sync"
)

type Base struct {
	Key   interface{}
	Value interface{}
}

type Store struct {
	db  []Base
	mux *sync.RWMutex
}

func NewStore(rwmux *sync.RWMutex) *Store {
	basetore := &Store{db: []Base{}, mux: rwmux}
	return basetore
}

func (s *Store) Append(key interface{}, value interface{}) {

	s.mux.Lock()
	s.db = append(s.db, Base{Key: key, Value: value})
	s.mux.Unlock()

}

func (s *Store) Get(key interface{}, v interface{}) {

	s.mux.RLock()
	var val interface{}
	for _, elem := range s.db {

		if elem.Key == key {
			val = elem.Value
			break
		}
	}
	v = &val
	s.mux.RUnlock()

}

func (s *Store) Get_index(key interface{}) (int, bool) {

	for index, elem := range s.db {
		if elem.Key == key {
			return index, true
		}
	}

	return 0, false
}

func (s *Store) Delete(key interface{}) {

	s.mux.Lock()

	i, ok := s.Get_index(key)
	if !ok {
		return
	}

	s.db[i] = s.db[len(s.db)-1]
	s.db = s.db[:len(s.db)-1]

	s.mux.Unlock()
}

func (s *Store) Update(key interface{}, value interface{}) {
	s.mux.Lock()
	i, ok := s.Get_index(key)
	fmt.Println("Get method ", ok)
	if !ok {
		return
	}

	m := s.db[i]
	m.Value = value
	s.mux.Unlock()

}

func main() {

	mux := &sync.RWMutex{}

	// Check Base structure
	//fmt.Println(Base{Key: "one", Value: 123})
	// Make new db structure
	one := NewStore(mux)

	// Check put method
	go one.Append("one", 123)
	go one.Append("two", 222)

	// Check get method
	var v interface{} = 123
	go one.Get("one", v)
	fmt.Println("Get method ", v)

	// Check update method
	go one.Update("one", 112)
	fmt.Println("Update ", one)

	// Check delete method
	//go one.Delete("one")
	//fmt.Println("delete ", one)

}
