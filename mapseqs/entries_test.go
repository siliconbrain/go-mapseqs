package mapseqs

import (
	"testing"

	"github.com/siliconbrain/go-seqs/seqs"
	"github.com/stretchr/testify/require"
)

func TestEntriesOf(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		require.Empty(t, seqs.ToSlice[MapEntry[string, any]](EntriesOf(map[string]any(nil))))
	})
	t.Run("empty map", func(t *testing.T) {
		require.Empty(t, seqs.ToSlice[MapEntry[string, any]](EntriesOf(map[string]any{})))
	})
	t.Run("non-empty map", func(t *testing.T) {
		require.ElementsMatch(t, []MapEntry[string, any]{{Key: "foo", Value: 42}, {Key: "bar", Value: false}}, seqs.ToSlice[MapEntry[string, any]](EntriesOf(map[string]any{"foo": 42, "bar": false})))
	})
}
