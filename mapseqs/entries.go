package mapseqs

import "github.com/siliconbrain/go-seqs/seqs"

func EntriesOf[M Mapping[K, V], K comparable, V any](m M) MapEntries[M, K, V] {
	return MapEntries[M, K, V]{
		Map: m,
	}
}

type MapEntries[M Mapping[K, V], K comparable, V any] struct {
	Map M
}

func (mes MapEntries[M, K, V]) ForEachUntil(fn func(MapEntry[K, V]) bool) {
	for k, v := range mes.Map {
		if fn(MapEntry[K, V]{
			Key:   k,
			Value: v,
		}) {
			return
		}
	}
}

func (mes MapEntries[M, K, V]) Len() int {
	return len(mes.Map)
}

type MapEntry[K any, V any] struct {
	Key   K
	Value V
}

var _ seqs.Seq[MapEntry[string, any]] = (*MapEntries[map[string]any, string, any])(nil)
var _ seqs.Lener = (*MapEntries[map[string]any, string, any])(nil)
