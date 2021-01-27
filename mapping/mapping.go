package mapping

import "github.com/robotlovesyou/mars"

// Map can be queried for obstacles
type Map struct {
	obstacles map[mars.Coordinate]bool
}

// New returns a new map
func New(obstacleList []mars.Coordinate) *Map {
	obstacles := make(map[mars.Coordinate]bool)
	for _, coordinate := range obstacleList {
		obstacles[coordinate] = true
	}
	return &Map{
		obstacles: obstacles,
	}
}

// HasObstacle returns true if a coordinate contains an obstacle
func (m *Map) HasObstacle(x, y int) bool {
	return m.obstacles[mars.Coordinate{X: x, Y: y}]
}
