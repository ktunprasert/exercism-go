package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	nodeMap := make(map[int]*Node)

	sort.Slice(records, func(i, j int) bool {
		if records[i].Parent == records[j].Parent {
			return records[i].ID < records[j].ID
		}
		return records[i].Parent < records[j].Parent
	})

	for _, r := range records {
		switch {
		case nodeMap[r.ID] != nil:
			return nil, errors.New("duplicate node")
		case r.ID >= len(records):
			return nil, errors.New("invalid ID")
		case r.Parent > r.ID:
			return nil, errors.New("parent ID greater than child ID")
		}

		nodeMap[r.ID] = &Node{ID: r.ID, Children: make([]*Node, 0)}

		if r.ID == 0 && r.Parent == 0 {
			continue
		}
		parent, ok := nodeMap[r.Parent]
		if !ok {
			return nil, errors.New("parent not found")
		}

		child := nodeMap[r.ID]
		if parent.ID == r.Parent && child.ID == r.Parent {
			return nil, errors.New("circular reference")
		}

		parent.Children = append(parent.Children, child)
	}

	return nodeMap[0], nil
}
