package maps

import (
	"testing"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/stretchr/testify/require"
)

func TestKeysOf(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		require.Empty(t, seqs.ToSlice[string](KeysOf(map[string]any(nil))))
	})
	t.Run("empty map", func(t *testing.T) {
		require.Empty(t, seqs.ToSlice[string](KeysOf(map[string]any{})))
	})
	t.Run("non-empty map", func(t *testing.T) {
		require.ElementsMatch(t, []string{"foo", "bar"}, seqs.ToSlice[string](KeysOf(map[string]any{"foo": 42, "bar": false})))
	})
}

func TestMapKeys_Has(t *testing.T) {
	keys := KeysOf(map[string]any{
		"foo": 42,
		"bar": false,
	})
	require.True(t, keys.Has("foo"))
	require.False(t, keys.Has("moose"))
}
