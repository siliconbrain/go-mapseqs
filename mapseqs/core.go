package mapseqs

type Mapping[K comparable, V any] interface {
	~map[K]V
}
