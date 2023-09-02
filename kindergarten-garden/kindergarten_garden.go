package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var plantsMap = map[string]string{
	"V": "violets",
	"R": "radishes",
	"C": "clover",
	"G": "grass",
}

// Define the Garden type here.
type Garden struct {
	gardenMap map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	gardenMap := make(map[string][]string)

	childrenCopy := make([]string, len(children))
	copy(childrenCopy, children)

	d := strings.Fields(diagram)

	switch {
	case len(d) != 2,
		len(d[0]) != len(d[1]),
		len(d[0])%2 != 0,
		len(d[1])%2 != 0,
		diagram[0] != '\n':

		return nil, errors.New("Invalid diagram")
	}

	sort.Strings(childrenCopy)
	prevName := ""
	for i, name := range childrenCopy {
		if prevName != "" && name == prevName {
			return nil, errors.New("Duplicate child name")
		}

		p1 := strings.Split(d[0][2*i:2*i+2], "")
		p2 := strings.Split(d[1][2*i:2*i+2], "")

		for _, p := range append(p1, p2...) {
			plantName, ok := plantsMap[string(p)]
			if !ok {
				return nil, errors.New("Invalid plant name")
			}

			gardenMap[name] = append(gardenMap[name], plantName)
		}

		prevName = name
	}

	g := &Garden{gardenMap}
	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	m, ok := g.gardenMap[child]
	return m, ok
}
