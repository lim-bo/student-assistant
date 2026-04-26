package routebuilder_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	routebuilder "github.com/student-assistant/pkg/route_builder"
)

func TestRouteBuilder(t *testing.T) {
	coords := []routebuilder.Coordinates{
		routebuilder.Coordinates([]float64{59.929033, 30.322471}),
		routebuilder.Coordinates([]float64{55.749153, 37.655177}),
		routebuilder.Coordinates([]float64{59.953713, 30.319469}),
	}

	route, err := routebuilder.Get(coords)
	assert.NoError(t, err)
	t.Log(route)
}
