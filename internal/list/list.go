package list

type List[T comparable] struct {
	allow *Set[T]
	deny  *Set[T]
}

func New[T comparable](allow []T, deny []T) *List[T] {
	return &List[T]{allow: NewSet(allow...), deny: NewSet(deny...)}
}

func (l *List[T]) Allow(v T) {
	l.allow.Add(v)
}

func (l *List[T]) AllowHas(v T) bool {
	return l.allow.Has(v)
}

func (l *List[T]) Deny(v T) {
	l.deny.Add(v)
}

func (l *List[T]) DenyHas(v T) bool {
	return l.deny.Has(v)
}

func (l *List[T]) Allows(v T) bool {
	return (l.allow.Empty() || l.allow.Has(v)) && !l.deny.Has(v)
}
