package lists

import (
	"errors"
	"fmt"
)

type List[T any] interface {
	Add(element T) error
	AddAt(index int, element T) error
	Get(index int) (T, error)
	Remove(index int) (T, error)
	Size() int
	Clear()
	Contains(element T, equals func(T, T) bool) bool
}

var size int

type ArrayList[T any] struct {
	elements []T
}

func NewArrayList[T any]() *ArrayList[T] {
	size = 0
	return &ArrayList[T]{
		elements: make([]T, 0),
	}
}

func (l *ArrayList[T]) Add(element T) {
	l.elements = append(l.elements, element)
	size++
}

func (l *ArrayList[T]) AddAt(index int, element T) error {
	if index < 0 || index > size {
		return errors.New("Index out of bounds")
	}
	for i := size - 1; i >= index; i-- {
		l.elements[i+1] = l.elements[i]
	}
	l.elements[index] = element
	size++
	return nil
}

func (l *ArrayList[T]) Get(index int) (T, error) {
	var zero T
	if index < 0 || index > size {
		return zero, errors.New("Index out of bounds")
	}

	return l.elements[index], nil
}

func (l *ArrayList[T]) Remove(index int) (T, error) {
	var zero T
	if index < 0 || index > size {
		return zero, errors.New("Index out of bounds")
	}
	element := l.elements[index]
	for i := index; i < size-1; i++ {
		l.elements[i] = l.elements[i+1]
	}
	size--
	return element, nil
}

func (l *ArrayList[T]) Size() int {
	return size
}

func (l *ArrayList[T]) Clear() {
	l.elements = make([]T, 0)
	size = 0
}
func (l *ArrayList[T]) Contains(element T, equals func(T, T) bool) bool {
	for i := 0; i < size; i++ {
		if equals(l.elements[i], element) {
			return true
		}
	}
	return false
}

func (l *ArrayList[T]) String() string {
	return fmt.Sprintf("%v", l.elements[:size])
}
