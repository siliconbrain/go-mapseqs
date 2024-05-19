package mapseqs

import (
	"testing"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/stretchr/testify/require"
)

func TestValuesOf(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		require.Empty(t, seqs.ToSlice(ValuesOf(map[string]any(nil))))
	})
	t.Run("empty map", func(t *testing.T) {
		require.Empty(t, seqs.ToSlice(ValuesOf(map[string]any{})))
	})
	t.Run("non-empty map", func(t *testing.T) {
		require.ElementsMatch(t, []any{42, false}, seqs.ToSlice(ValuesOf(map[string]any{"foo": 42, "bar": false})))
	})
}
