package routebuilder

import (
	"fmt"
	"strings"
)

type Coordinates [2]float64

func (c *Coordinates) String() string {
	return fmt.Sprintf("%f,%f", c[0], c[1])
}

const mapSourceYandex = "https://yandex.ru/maps/"

func Get(coords []Coordinates) (string, error) {
	if len(coords) < 2 {
		return "", fmt.Errorf("Cannot build route with less than 2 points")
	}
	url := mapSourceYandex + "?"

	coordsStrSlice := make([]string, 0, len(coords))
	for _, c := range coords {
		coordsStrSlice = append(coordsStrSlice, c.String())
	}
	url += fmt.Sprintf("rtext=%s", strings.Join(coordsStrSlice, "~"))
	return url, nil
}
