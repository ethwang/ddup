package setcode

import "sync"

type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}

func NewSet() *Set {

	return &Set{
		m: make(map[interface{}]struct{}),
	}
}

func (s *Set) Add(ks ...interface{}) bool {
	if s == nil || s.m == nil {
		return false
	}
	s.Lock()
	defer s.Unlock()

	for _, v := range ks {
		s.m[v] = struct{}{}
	}

	return true

}
func (s *Set) Remove(ks ...interface{}) bool {

	if s == nil || s.m == nil {
		return false
	}
	s.Lock()
	defer s.Unlock()

	for _, v := range ks {
		if s.Has(v) {
			delete(s.m, v)
		}
	}
	return true
}
func (s *Set) Has(k interface{}) bool {
	if s == nil || s.m == nil {
		return false
	}

	s.RLock()
	defer s.RUnlock()

	if _, ok := s.m[k]; ok {
		return true
	}

	return false
}
func (s *Set) All() []interface{} {

	if s == nil || s.m == nil {
		return nil
	}
	var list []interface{}

	s.RLock()
	defer s.RUnlock()

	for _, v := range s.m {
		list = append(list, v)
	}

	return list
}
