package mapping_test

import (
	"testing"

	"github.com/robotlovesyou/mars/position"

	"github.com/robotlovesyou/mars/mapping"
	"github.com/stretchr/testify/require"
)

var obstacles = []position.Coordinate{{1, 4}, {3, 5}, {7, 4}}

func TestHasObstacleReturnsFalseForASafeCoordinate(t *testing.T) {
	r := require.New(t)
	m := mapping.New(obstacles)
	r.False(m.HasObstacle(1, 1))
}

func TestHasObstacleReturnsTrueForDangerousCoordinate(t *testing.T) {
	r := require.New(t)
	m := mapping.New(obstacles)
	r.True(m.HasObstacle(1, 4))
}
