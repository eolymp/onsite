package list

import "sync"

type Set[T comparable] struct {
	lock   sync.RWMutex
	values map[T]bool
}

func NewSet[T comparable](values ...T) *Set[T] {
	s := &Set[T]{values: map[T]bool{}}
	s.Add(values...)

	return s
}

func (s *Set[T]) Add(values ...T) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, value := range values {
		s.values[value] = true
	}
}

func (s *Set[T]) Has(value T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, ok := s.values[value]
	return ok
}

func (s *Set[T]) Empty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.values) == 0
}
