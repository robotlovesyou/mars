package mapping

import "github.com/robotlovesyou/mars/position"

// Map can be queried for obstacles
type Map struct {
	obstacles map[position.Coordinate]bool
}

// NewPosition returns a new map
func New(obstacleList []position.Coordinate) *Map {
	obstacles := make(map[position.Coordinate]bool)
	for _, coordinate := range obstacleList {
		obstacles[coordinate] = true
	}
	return &Map{
		obstacles: obstacles,
	}
}

// HasObstacle returns true if a coordinate contains an obstacle
func (m *Map) HasObstacle(coord position.Coordinate) bool {
	return m.obstacles[coord]
}
