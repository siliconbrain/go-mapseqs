package mapseqs

import "github.com/siliconbrain/go-seqs/seqs"

func EntriesOf[
	Map Mapping[Key, Value],
	Key comparable,
	Value any,
](
	m Map,
) MapEntries[
	Map,
	Key,
	Value,
	MapEntry[Key, Value],
	func(Key, Value) MapEntry[Key, Value],
] {
	return EntriesOfWith(m, MapEntryFrom)
}

func EntriesOfWith[
	Map Mapping[Key, Value],
	Key comparable,
	Value any,
	Entry any,
	Pack ~func(Key, Value) Entry,
](m Map, pack Pack) MapEntries[Map, Key, Value, Entry, Pack] {
	return MapEntries[Map, Key, Value, Entry, Pack]{
		Map:  m,
		Pack: pack,
	}
}

type MapEntries[
	Map Mapping[Key, Value],
	Key comparable,
	Value any,
	Entry any,
	Pack ~func(Key, Value) Entry,
] struct {
	Map  Map
	Pack Pack
}

func (mes MapEntries[_, K, V, E, _]) ForEachUntil(yield func(E) bool) {
	for k, v := range mes.Map {
		if yield(mes.Pack(k, v)) {
			return
		}
	}
}

func (mes MapEntries[_, _, _, _, _]) Len() int {
	return len(mes.Map)
}

func MapEntryFrom[Key, Value any](key Key, value Value) MapEntry[Key, Value] {
	return MapEntry[Key, Value]{
		Key:   key,
		Value: value,
	}
}

type MapEntry[Key, Value any] struct {
	Key   Key
	Value Value
}

var _ seqs.Seq[any] = (*MapEntries[map[string]any, string, any, any, func(string, any) any])(nil)
var _ seqs.Lener = (*MapEntries[map[string]any, string, any, any, func(string, any) any])(nil)
