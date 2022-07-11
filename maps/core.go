package maps

type Mapping[K comparable, V any] interface {
	~map[K]V
}
