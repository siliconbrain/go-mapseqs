package mapseqs

import "github.com/siliconbrain/go-seqs/seqs"

func KeysOf[M Mapping[K, V], K comparable, V any](m M) MapKeys[M, K, V] {
	return MapKeys[M, K, V]{
		Map: m,
	}
}

type MapKeys[M Mapping[K, V], K comparable, V any] struct {
	Map M
}

func (mks MapKeys[M, K, V]) ForEachUntil(yield func(K) bool) {
	for k := range mks.Map {
		if yield(k) {
			return
		}
	}
}

func (mks MapKeys[M, K, V]) Has(k K) bool {
	_, has := mks.Map[k]
	return has
}

func (mks MapKeys[M, K, V]) Len() int {
	return len(mks.Map)
}

var _ seqs.Seq[string] = (*MapKeys[map[string]any, string, any])(nil)
var _ seqs.Lener = (*MapKeys[map[string]any, string, any])(nil)
