package gmvvm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestViewModel(t *testing.T) {
	var value int = 3
	model := NewModel[int](&value)
	view := NewView[int]()
	WatchModel[int](model, view)
	require.Equal(t, 0, view.Value)
	model.UpdateViews()
	require.Equal(t, 3, view.Value)
	view.Value = 4
	view.UpdateModel()
	require.Equal(t, 4, value)
	require.Equal(t, 4, *model.Value)
	model.UpdateWith(func(orgin int) int {
		return orgin + 2
	})
	require.Equal(t, 6, value)
	require.Equal(t, 6, view.Value)
}
