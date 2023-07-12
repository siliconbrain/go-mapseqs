package mapseqs

import "github.com/siliconbrain/go-seqs/seqs"

func ValuesOf[M Mapping[K, V], K comparable, V any](m M) MapValues[M, K, V] {
	return MapValues[M, K, V]{
		Map: m,
	}
}

type MapValues[M Mapping[K, V], K comparable, V any] struct {
	Map M
}

func (mks MapValues[M, K, V]) ForEachUntil(fn func(V) bool) {
	for _, v := range mks.Map {
		if fn(v) {
			return
		}
	}
}

func (mks MapValues[M, K, V]) Len() int {
	return len(mks.Map)
}

var _ seqs.Seq[any] = (*MapValues[map[string]any, string, any])(nil)
var _ seqs.Lener = (*MapValues[map[string]any, string, any])(nil)
