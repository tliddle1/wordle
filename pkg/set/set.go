package set

type Set[T comparable] map[T]struct{}

func (this Set[T]) Add(value T) {
	this[value] = struct{}{}
}

func (this Set[T]) Remove(value T) {
	delete(this, value)
}

func (this Set[T]) Contains(value T) bool {
	_, ok := this[value]
	return ok
}
