package config

type Coordinate struct {
	Row, Col int
}

// Define multiple patterns. Each pattern is a slice of coordinates.
var Patterns = [][]Coordinate{
	// Middle horizontal line
	{
		{1, 0},
		{1, 1},
		{1, 2},
		{1, 3},
	},
	// Add more patterns as needed
}
